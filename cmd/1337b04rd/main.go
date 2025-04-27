package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"1337B04RD/config"
	utils "1337B04RD/helper"
	adapter_db "1337B04RD/internal/adapter/db"
	adapter_http "1337B04RD/internal/adapter/http"
	"1337B04RD/internal/adapter/postgres"
	"1337B04RD/internal/app/common/logger"
	domain_port "1337B04RD/internal/domain/port"
	"1337B04RD/internal/domain/service"
)

func main() {
	port := flag.Int("port", 8080, "Listening port number")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		utils.Helper()
		return
	}

	// Load config and init logger
	cfg := config.Load()
	logger.Init(cfg.AppEnv)

	// connection postgreSQL
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Error("failed to connect to DB", "err", err)
		return
	}
	defer db.Close()

	logger.Info("connected to PostgreSQL", "host", cfg.DB.Host, "db", cfg.DB.Name)

	dbAdapter := adapter_db.NewPostgresRepository(db)

	// rickMortyAPI := initRickMortyAPI()

	// Инициализация репозиториев
	var postRepo domain_port.PostRepository = dbAdapter
	var commentRepo domain_port.CommentRepository = dbAdapter
	var sessionRepo domain_port.SessionRepository = dbAdapter

	// Инициализация сервисов
	postService := service.NewPostService(postRepo, commentRepo, s3, logger)
	// rickMortyAPI добавить в параметры :
	sessionService := service.NewSessionService(sessionRepo, logger)

	// Инициализация обработчиков
	handler := adapter_http.NewHandler(postService, sessionService)
	// Настройка маршрутизации
	router := adapter_http.SetupRoutes(handler, sessionService)
	// Другие маршруты

	addr := fmt.Sprintf(":%d", *port)
	// Запуск сервера
	logger.Info("🚀Starting server on port " + addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
