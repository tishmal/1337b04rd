package db

import "database/sql"

// internal/adapter/db/postgres.go
type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
	// Инициализация подключения к БД
}

// Реализация методов PostRepository, CommentRepository, SessionRepository
