package cmd

import (
	"1337B04RD/internal/domain/port"
	"1337B04RD/internal/domain/service"
	"1337B04RD/utils"
	"flag"
	"fmt"
	"http"
	"log/slog"
	"os"
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
	postService := service.NewPostService(postRepo, commentRepo, s3)
	// sessionService := service.NewSessionService(sessionRepo, rickMortyAPI)

	// Инициализация обработчиков
	// handler := http.NewHandler(postService, sessionService)
	handler := http.NewHandler(postService)
	// Настройка маршрутизации
	router := http.NewServeMux()
	router.HandleFunc("/", handler.CatalogHandler)
	// Другие маршруты

	addr := fmt.Sprintf(":%d", *portS)
	// Запуск сервера
	logger.Info("Starting server on port " + addr)
	err := http.ListenAndServe(":"+addr, router)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
