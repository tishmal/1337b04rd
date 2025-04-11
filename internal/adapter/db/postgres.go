package db

import (
	"database/sql"
)

// internal/adapter/db/postgres.go
type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(_db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: _db}
}

// Реализация методов PostRepository, CommentRepository, SessionRepository
