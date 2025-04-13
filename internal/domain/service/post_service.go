package service

import (
	"1337B04RD/internal/domain/entity"
	"1337B04RD/internal/domain/port"
)

type PostService struct {
	postRepo     port.PostRepository
	commentRepo  port.CommentRepository
	imageStorage port.ImageStorage
}

// Инициализация сервиса
func NewPostService(postRepo port.PostRepository, commentRepo port.CommentRepository, imageStorage port.ImageStorage) *PostService {
	return &PostService{
		postRepo:     postRepo,
		commentRepo:  commentRepo,
		imageStorage: imageStorage}
}

// // Методы для работы с постами
// (title, content string, image []byte, userID, userName, avatarURL string)
func (s *PostService) CreatePost(title, content, sessionID string) (*entity.Post, error) {
	// Логика создания поста

}

func (s *PostService) UploadImage(buf []byte, filename string) (string, error) {

}
