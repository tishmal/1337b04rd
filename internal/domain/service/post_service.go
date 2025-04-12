package service

import (
	"1337B04RD/internal/domain/port"
)

// internal/domain/service/post_service.go
type PostService struct {
	postRepo     port.PostRepository
	commentRepo  port.CommentRepository
	imageStorage port.ImageStorage
}

func NewPostService(postRepo port.PostRepository, commentRepo port.CommentRepository, imageStorage port.ImageStorage) *PostService {
	// Инициализация сервиса
	return &PostService{}
	// заглушка
}

// // Методы для работы с постами
// func (s *PostService) CreatePost(title, content string, image []byte, userID, userName, avatarURL string) (*entity.Post, error) {
// 	// Логика создания поста
// }
