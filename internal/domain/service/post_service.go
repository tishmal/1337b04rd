package service

import (
	"1337B04RD/internal/domain/entity"
	"1337B04RD/internal/domain/port"
	"time"

	"github.com/google/uuid"
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
// Логика создания поста
func (s *PostService) CreatePost(title, content, sessionID string) (*entity.Post, error) {
	var newPost entity.Post
	newPost.ID = uuid.New().String()
	newPost.CreatedAt = time.Now()
	newPost.Content = content
	newPost.Title = title

	err := s.postRepo.Create(&newPost)
	if err != nil {
		return &entity.Post{}, err
	}

	return &newPost, nil
}

// func (s *PostService) UploadImage(buf []byte, filename string) (string, error) {

// }
