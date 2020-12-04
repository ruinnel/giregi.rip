package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	_user "github.com/ruinnel/giregi.rip-server/feature/user/delivery/http"
	_userService "github.com/ruinnel/giregi.rip-server/feature/user/service"
	"github.com/ruinnel/giregi.rip-server/http/middleware"
	"github.com/ruinnel/giregi.rip-server/repository"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"

	_archive "github.com/ruinnel/giregi.rip-server/feature/archive/delivery/http"
	_archiveCache "github.com/ruinnel/giregi.rip-server/feature/archive/repository/redis"
	_archiveService "github.com/ruinnel/giregi.rip-server/feature/archive/service"

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

	err := repository.Use(config)
	if err != nil {
		logger.Fatal(err)
	}

	cache := common.OpenRedis(config.Redis)

	defer repository.Disconnect()

	logger.Printf("mode - %v", mode)
	switch mode {
	case "worker":
		runWorker(config, cache)
	default:
		runServer(config, cache)
	}
}

func runServer(config *common.Config, cache *redis.Client) {
	userRepository := repository.User()
	tokenRepository := repository.Token()
	archiveRepository := repository.Archive()
	siteRepository := repository.Site()
	webPageRepository := repository.WebPage()
	tagRepository := repository.Tag()
	archiveCache := _archiveCache.NewArchiveCache(cache)

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

func runWorker(config *common.Config, cache *redis.Client) {
	logger := common.GetLogger()
	archiveRepository := repository.Archive()
	siteRepository := repository.Site()
	webPageRepository := repository.WebPage()
	tagRepository := repository.Tag()
	archiveCache := _archiveCache.NewArchiveCache(cache)

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
	err = service.CheckProgress(context.Background(), archive)
	if err != nil {
		logger.Printf("archive fail: %v", err)
	}
}
