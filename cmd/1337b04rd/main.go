package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"1337B04RD/internal/adapter/db"
	adapter_http "1337B04RD/internal/adapter/http"
	"1337B04RD/internal/domain/port"
	"1337B04RD/internal/domain/service"
	"1337B04RD/utils"

	_ "github.com/lib/pq"
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

	// Инициализация адаптеров (используем функции из init.go)
	dbAdapter := initDatabase(logger)
	s3 := initS3Storage(logger)

	// Запуск миграций, если AUTO_MIGRATE=true
	if shouldRunMigrations() {
		if err := db.RunMigrations(dbAdapter.GetDB(), logger); err != nil {
			logger.Error("Failed to run migrations", "error", err)
			os.Exit(1)
		}
		logger.Info("Migrations completed successfully")
	}

	// rickMortyAPI := initRickMortyAPI()

	// Инициализация репозиториев
	var postRepo port.PostRepository = dbAdapter
	var commentRepo port.CommentRepository = dbAdapter
	// var sessionRepo port.SessionRepository = dbAdapter

	// Инициализация сервисов
	postService := service.NewPostService(postRepo, commentRepo, s3)
	// sessionService := service.NewSessionService(sessionRepo, rickMortyAPI)

	// Инициализация обработчиков
	handler := adapter_http.NewHandler(postService)
	// handler := adapter_http.NewHandler(postService, sessionService)

	// Настройка маршрутизации
	router := adapter_http.SetupRoutes(handler)

	// Запуск сервера
	addr := fmt.Sprintf(":%d", *portS)
	logger.Info("Starting server on port " + addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

// shouldRunMigrations проверяет, нужно ли запускать миграции
func shouldRunMigrations() bool {
	return getEnvAsBool("AUTO_MIGRATE", true)
}

// getEnvAsBool получает булеву переменную окружения или возвращает значение по умолчанию
func getEnvAsBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
