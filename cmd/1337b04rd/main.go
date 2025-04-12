package cmd

import (
	"log/slog"
	"os"
)

func main() {
	// Парсинг аргументов командной строки

	// Инициализация логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Инициализация адаптеров
	db := initDatabase(logger)
	s3 := initS3Storage(logger)
	// rickMortyAPI := initRickMortyAPI()

	// Инициализация репозиториев
	postRepo := db.NewPostRepository()
	commentRepo := db.NewCommentRepository()
	sessionRepo := db.NewSessionRepository()

	// Инициализация сервисов
	postService := service.NewPostService(postRepo, commentRepo, s3)
	sessionService := service.NewSessionService(sessionRepo, rickMortyAPI)

	// Инициализация обработчиков
	handler := http.NewHandler(postService, sessionService)

	// Настройка маршрутизации
	router := http.NewServeMux()
	router.HandleFunc("/", handler.CatalogHandler)
	// Другие маршруты

	// Запуск сервера
	logger.Info("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
