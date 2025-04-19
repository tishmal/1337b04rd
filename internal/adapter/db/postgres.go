package db

import (
	"context"
	"database/sql"

	"1337B04RD/internal/domain/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(_db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: _db}
}

// internal/adapter/db/postgres_repository.go
// Добавьте этот метод к вашему существующему PostgresRepository

// GetDB возвращает экземпляр sql.DB для использования в миграциях
func (r *PostgresRepository) GetDB() *sql.DB {
	return r.db
}

// Реализация методов PostRepository, CommentRepository, SessionRepository

// Реализация интерфейса PostRepository
func (r *PostgresRepository) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	query := `INSERT INTO posts (post_id, title, content, user_avatar_id, session_id, user_name)
	          VALUES ($1, $2, $3, $4, $5, $6)
	          RETURNING id, post_id, title, content, user_avatar_id, user_name, session_id, created_at`

	row := r.db.QueryRowContext(ctx, query,
		post.PostID,
		post.Title,
		post.Content,
		post.UserAvatarID,
		post.SessionID,
		post.UserName,
	)

	err := row.Scan(&post.ID, &post.PostID, &post.Title, &post.Content, &post.UserAvatarID, &post.UserName, &post.SessionID, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}
