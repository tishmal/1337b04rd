package port

import (
	"1337B04RD/internal/domain/entity"
	"context"
)

type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	// Get(id string) (*entity.Post, error)
	// GetAll() ([]*entity.Post, error)
	// GetArchived() ([]*entity.Post, error)
	// Update(post *entity.Post) error
	// Delete(id string) error
	// // Другие методы...
}

type CommentRepository interface {
	// Create(comment *entity.Comment) error
	// GetByPostID(postID string) ([]*entity.Comment, error)
	// // Другие методы...
}

type SessionRepository interface {
	Save(session *entity.Session) error
	Get(sessionID string) (*entity.Session, error)
	Delete(sessionID string) error
}
