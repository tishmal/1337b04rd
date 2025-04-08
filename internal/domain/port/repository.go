package port

import "1337B04RD/internal/domain/entity"

type PostRepository interface {
	Create(post *entity.Post) error
	Get(id string) (*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	GetArchived() ([]*entity.Post, error)
	Update(post *entity.Post) error
	Delete(id string) error
	// Другие методы...
}

type CommentRepository interface {
	Create(comment *entity.Comment) error
	GetByPostID(postID string) ([]*entity.Comment, error)
	// Другие методы...
}

type SessionRepository interface {
	Create(session *entity.Session) error
	Get(id string) (*entity.Session, error)
	Delete(id string) error
	// Другие методы...
}
