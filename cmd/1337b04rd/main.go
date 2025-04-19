package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	adapter_http "1337B04RD/internal/adapter/http"
	"1337B04RD/internal/domain/port"
	"1337B04RD/internal/domain/service"
	"1337B04RD/utils"
)

func main() {
	// Парсинг аргументов командной строки
	portS := flag.Int("port", 8080, "Listening port number")
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
	var postRepo port.PostRepository = dbAdapter
	var commentRepo port.CommentRepository = dbAdapter
	// var sessionRepo port.SessionRepository = dbAdapter

	// Инициализация сервисов
	postService := service.NewPostService(postRepo, commentRepo, s3, logger)
	// sessionService := service.NewSessionService(sessionRepo, rickMortyAPI)

	// Инициализация обработчиков
	handler := adapter_http.NewHandler(postService)
	// handler := adapter_http.NewHandler(postService, sessionService)
	// Настройка маршрутизации
	router := adapter_http.SetupRoutes(handler)
	// Другие маршруты

	addr := fmt.Sprintf(":%d", *portS)
	// Запуск сервера
	logger.Info("🚀Starting server on port " + addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
