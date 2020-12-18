package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alitto/pond"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	_archive "github.com/ruinnel/giregi.rip-server/feature/archive/delivery/http"
	_archiveService "github.com/ruinnel/giregi.rip-server/feature/archive/service"
	_user "github.com/ruinnel/giregi.rip-server/feature/user/delivery/http"
	_userService "github.com/ruinnel/giregi.rip-server/feature/user/service"
	"github.com/ruinnel/giregi.rip-server/http/middleware"
	"github.com/ruinnel/giregi.rip-server/queue"
	"github.com/ruinnel/giregi.rip-server/repository"
	"os"
	"os/signal"
	"syscall"

	"log"
	"net/http"
	"time"
)

var (
	version  = ""
	platform = ""
)

//goland:noinspection GoBoolExpressions
func main() {
	logger := common.GetLogger()

	if !(platform == string(common.PLATFORM_SERVER) || platform == string(common.PLATFORM_DESKTOP)) {
		panic("unknown mode")
	}

	logger.Printf("version - %v, platform - %v", version, platform)

	mode := "server"
	yamlFile := "./config.yaml"
	if platform == string(common.PLATFORM_SERVER) {
		flag.StringVar(&mode, "mode", "server", "mode - server or worker")
	}
	flag.StringVar(&yamlFile, "config", "./config.yaml", "require config file")
	flag.Parse()
	logger.Printf("service key json - %s\n", yamlFile)
	if len(yamlFile) == 0 {
		panic("error: config json file not found.")
	}

	config := common.InitConfig(yamlFile)
	config.Platform = common.Platform(platform)

	err := repository.Use(config)
	if err != nil {
		logger.Fatal(err)
	}
	defer repository.Disconnect()

	q := queue.NewQueue(config)
	defer q.Close()

	if config.Platform == common.PLATFORM_SERVER {
		logger.Printf("mode - %v", mode)
		switch mode {
		case "worker":
			runWorker(config, q)
		case "server":
			runServer(config, q)
		}
	} else {
		logger.Printf("platform - %v", config.Platform)
		go runWorker(config, q)
		runServer(config, q)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			os.Exit(0)
		default:
			fmt.Printf("Unknown signal(%d)\n", s)
		}
	}
}

func runServer(config *common.Config, q queue.Queue) {
	userRepository := repository.User()
	tokenRepository := repository.Token()
	archiveRepository := repository.Archive()
	siteRepository := repository.Site()
	webPageRepository := repository.WebPage()
	tagRepository := repository.Tag()
	archiveCache := repository.ArchiveCache()

	userService := _userService.NewUserService(
		userRepository, tokenRepository, tagRepository, archiveRepository,
	)
	archiveService := _archiveService.NewArchiveService(
		archiveRepository, archiveCache,
		siteRepository, webPageRepository,
		tagRepository, q,
	)

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware(config, userService))
	router.Use(middleware.UrlDecodeMiddleware())

	prefixRouter := router.PathPrefix(config.Server.ContextPath).Subrouter()
	_user.User(prefixRouter.PathPrefix("/users").Subrouter(), userService)
	_archive.Archive(prefixRouter.PathPrefix("/archives").Subrouter(), archiveService)

	server := &http.Server{
		Handler: cors.AllowAll().Handler(router),
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
	}

	log.Printf("start server - %v\n", fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))

	log.Fatal(server.ListenAndServe())
}

func runWorker(config *common.Config, q queue.Queue) {
	logger := common.GetLogger()

	pool := pond.New(config.WorkerSize, config.WorkerSize*10)
	archiveRepository := repository.Archive()
	siteRepository := repository.Site()
	webPageRepository := repository.WebPage()
	tagRepository := repository.Tag()
	archiveCache := repository.ArchiveCache()

	archiveService := _archiveService.NewArchiveService(
		archiveRepository, archiveCache,
		siteRepository, webPageRepository,
		tagRepository, q,
	)

	ch, err := q.Channel()
	if err != nil {
		logger.Panicf("init queue fail - %v", err)
	}

	for archive := range ch {
		logger.Printf("check archive progress - %v", archive)
		pool.Submit(func() {
			checkArchiveProgress(archive, archiveService)
		})
	}
}

func checkArchiveProgress(archive *domain.Archive, service domain.ArchiveService) {
	logger := common.GetLogger()
	err := service.CheckProgress(context.Background(), archive)
	if err != nil {
		logger.Printf("archive fail: %v", err)
	}
}
