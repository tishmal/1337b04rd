package cmd

import (
	database "1337B04RD/internal/adapter/db"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
)

func initDatabase() *database.PostgresRepository {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "1337b04rd")

	// Такой формат привычный для lib/pq, всё чётко.
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger := slog.Default()
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

	return database.NewPostgresRepository(db)
}

// Функция для получения переменных окружения с дефолтным значением
func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
