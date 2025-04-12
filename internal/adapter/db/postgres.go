package db

import (
	"1337B04RD/internal/domain/entity"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(_db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: _db}
}

// Реализация методов PostRepository, CommentRepository, SessionRepository

// Реализация интерфейса PostRepository
func (p *PostgresRepository) Create(post *entity.Post) error {
	// реализация
	var err error
	return err
}
