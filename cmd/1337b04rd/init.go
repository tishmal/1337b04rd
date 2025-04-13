package main

import (
	dbadapter "1337B04RD/internal/adapter/db"
	storageS3 "1337B04RD/internal/adapter/storage"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func initDatabase(logger *slog.Logger) *dbadapter.PostgresRepository {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "1337b04rd")

	// Такой формат привычный для lib/pq, всё чётко.
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.Info("Connecting to database", "host", host, "port", port, "database", dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Это обязательная проверка, часто её пропускают
	err = db.Ping()
	if err != nil {
		logger.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	logger.Info("Successfully connected to database")

	return dbadapter.NewPostgresRepository(db)
}

func initS3Storage(logger *slog.Logger) *storageS3.S3Storage {
	endpoint := getEnv("S3_ENDPOINT", "localhost:9000")
	accessKey := getEnv("S3_ACCESS_KEY", "minioadmin")
	secretKey := getEnv("S3_SECRET_KEY", "minioadmin")
	useSSL := getEnv("S3_USE_SSL", "false") == "true"
	postBucket := getEnv("S3_POST_BUCKET", "posts")
	commentBucket := getEnv("S3_COMMENT_BUCKET", "comments")

	logger.Info("Initializing S3 storage", "endpoint", endpoint)

	// Инициализация клиента S3
	storage, err := storageS3.NewS3Storage(endpoint, accessKey, secretKey, useSSL, postBucket, commentBucket, logger)
	if err != nil {
		logger.Error("Failed to initialize S3 storage", "error", err)
		os.Exit(1)
	}

	logger.Info("Successfully initialized S3 storage")
	return storage
}

// Функция для получения переменных окружения с дефолтным значением
func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
