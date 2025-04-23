package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	utils "1337B04RD/helper"
	adapter_http "1337B04RD/internal/adapter/http"
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

	// Инициализация логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Инициализация адаптеров
	dbAdapter := initDatabase(logger)
	s3 := initS3Storage(logger)
	// rickMortyAPI := initRickMortyAPI()

	// Инициализация репозиториев
	var postRepo domain_port.PostRepository = dbAdapter
	var commentRepo domain_port.CommentRepository = dbAdapter
	var sessionRepo domain_port.SessionRepository = dbAdapter

	// Инициализация сервисов
	postService := service.NewPostService(postRepo, commentRepo, s3, logger)
	// rickMortyAPI добавить в параметры :
	sessionService := service.NewSessionService(sessionRepo)

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
