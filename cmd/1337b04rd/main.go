package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"1337b04rd/config"
	utils "1337b04rd/helper"
	adapter_db "1337b04rd/internal/adapter/db"
	adapter_http "1337b04rd/internal/adapter/http"
	"1337b04rd/internal/adapter/postgres"
	"1337b04rd/internal/adapter/s3"
	"1337b04rd/internal/app/common/logger"
	domain_port "1337b04rd/internal/domain/port"
	"1337b04rd/internal/domain/services"
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

	//httpClient := &http.Client{}

	// init repositories
	var postRepo domain_port.PostRepository = dbAdapter
	var commentRepo domain_port.CommentRepository = dbAdapter
	var sessionRepo domain_port.SessionRepository = dbAdapter

	// S3 clients for threads and comments
	s3ThreadsClient := s3.NewS3Client(cfg.S3.Endpoint, cfg.S3.BucketThreads)
	s3CommentsClient := s3.NewS3Client(cfg.S3.Endpoint, cfg.S3.BucketComments)

	// Services
	// avatar
	sessionService := services.NewSessionService(sessionRepo)

	postS3Adapter := s3.NewAdapter(s3ThreadsClient)
	commentS3Adapter := s3.NewAdapter(s3CommentsClient)

	postService := services.NewPostService(postRepo, commentRepo, postS3Adapter)

	// Инициализация обработчиков
	handler := adapter_http.NewHandler(postService, sessionService)
	// Настройка маршрутизации
	router := adapter_http.SetupRoutes(handler, sessionService)
	// Другие маршруты

	addr := fmt.Sprintf(":%d", *port)
	// Запуск сервера
	logger.Info("🚀Starting server on port " + addr)
	errS := http.ListenAndServe(addr, router)
	if errS != nil {
		logger.Error("Failed to start server", "error", errS)
		os.Exit(1)
	}
}
