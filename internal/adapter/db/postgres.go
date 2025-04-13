package db

import (
	"1337B04RD/internal/domain/entity"
	"context"
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
func (p *PostgresRepository) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	query := `INSERT INTO posts (id, title, content)
          VALUES ($1, $2, $3)
          RETURNING id, title, content, created_at`

	row := p.db.QueryRowContext(ctx, query, post.ID, post.Title, post.Content)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return &entity.Post{}, err
	}
	return post, nil
}
