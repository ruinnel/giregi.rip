package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/http/middleware"
	_user "github.com/ruinnel/giregi.rip-server/user/delivery/http"
	_userRepository "github.com/ruinnel/giregi.rip-server/user/repository/mysql"
	_userService "github.com/ruinnel/giregi.rip-server/user/service"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"

	_archive "github.com/ruinnel/giregi.rip-server/archive/delivery/http"
	_archiveRepository "github.com/ruinnel/giregi.rip-server/archive/repository/mysql"
	_archiveCache "github.com/ruinnel/giregi.rip-server/archive/repository/redis"
	_archiveService "github.com/ruinnel/giregi.rip-server/archive/service"

	_siteRepository "github.com/ruinnel/giregi.rip-server/site/repository/mysql"
	_tagRepository "github.com/ruinnel/giregi.rip-server/tag/repository/mysql"
	_tokenRepository "github.com/ruinnel/giregi.rip-server/token/repository/mysql"
	_webPageRepository "github.com/ruinnel/giregi.rip-server/webpage/repository/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"net/http"
	"time"
)

func main() {
	logger := common.GetLogger()
	mode := "server"
	yamlFile := "./config.yaml"
	flag.StringVar(&mode, "mode", "server", "mode - server or worker")
	flag.StringVar(&yamlFile, "config", "./config.yaml", "require config file")
	flag.Parse()
	logger.Printf("service key json - %s\n", yamlFile)
	if len(yamlFile) == 0 {
		panic("error: credential json file not found.")
	}

	config := common.InitConfig(yamlFile)

	db := common.OpenDatabase(config.Database)
	cache := common.OpenRedis(config.Redis)

	migrateDatabase(config, db)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	boil.SetDB(db)
	// boil.DebugMode = true

	logger.Printf("mode - %v", mode)
	switch mode {
	case "worker":
		runWorker(config, db, cache)
	default:
		runServer(config, db, cache)
	}
}

func migrateDatabase(config *common.Config, db *sql.DB) {
	logger := common.GetLogger()
	source := migrate.FileMigrationSource{
		Dir: config.SQLMigrateSourcePath,
	}
	applyCount, err := migrate.Exec(db, "mysql", source, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("error: migration source(%s) not found. - %v", config.SQLMigrateSourcePath, err))
	}
	logger.Printf("migrate complete - %v", applyCount)
}

func runServer(config *common.Config, db *sql.DB, cache *redis.Client) {
	userRepository := _userRepository.NewUserRepository(db)
	tokenRepository := _tokenRepository.NewTokenRepository(db)
	archiveRepository := _archiveRepository.NewArchiveRepository(db)
	archiveCache := _archiveCache.NewArchiveCache(cache)
	siteRepository := _siteRepository.NewSiteRepository(db)
	webPageRepository := _webPageRepository.NewWebPageRepository(db)
	tagRepository := _tagRepository.NewTagRepository(db)

	userService := _userService.NewUserService(
		userRepository, tokenRepository, tagRepository, archiveRepository,
	)
	archiveService := _archiveService.NewArchiveService(
		archiveRepository, archiveCache,
		siteRepository, webPageRepository,
		tagRepository, config.RabbitMQ,
	)

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware(config, userService))
	router.Use(middleware.UrlDecodeMiddleware())

	prefixRouter := router.PathPrefix(config.Server.ContextPath)
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

func runWorker(config *common.Config, db *sql.DB, cache *redis.Client) {
	logger := common.GetLogger()
	archiveRepository := _archiveRepository.NewArchiveRepository(db)
	archiveCache := _archiveCache.NewArchiveCache(cache)
	siteRepository := _siteRepository.NewSiteRepository(db)
	webPageRepository := _webPageRepository.NewWebPageRepository(db)
	tagRepository := _tagRepository.NewTagRepository(db)

	archiveService := _archiveService.NewArchiveService(
		archiveRepository, archiveCache,
		siteRepository, webPageRepository,
		tagRepository, config.RabbitMQ,
	)

	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.RabbitMQ.Username, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port)
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		logger.Printf("amqp - %v", rabbitMQUrl)
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(config.RabbitMQ.Queue, false, false, false, false, nil)
	if err != nil {
		logger.Fatalf("connect to rabbitMQ(amqp) fail: %v", err)
	}

	messageChannel, err := ch.Consume(config.RabbitMQ.Queue, "", true, false, false, false, nil)
	if err != nil {
		logger.Fatalf("consume message fail: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for msg := range messageChannel {
			go processArchive(msg, archiveService)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

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

func processArchive(msg amqp.Delivery, service domain.ArchiveService) {
	logger := common.GetLogger()
	log.Printf("Received a message: %s", msg.Body)
	archive := new(domain.Archive)
	err := json.Unmarshal(msg.Body, archive)
	if err != nil {
		logger.Printf("unmarshal fail: %v", err)
		return
	}
	err = service.ProcessArchive(context.Background(), archive)
	if err != nil {
		logger.Printf("archive fail: %v", err)
	}
}
