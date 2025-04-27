package service

import (
	"1337B04RD/internal/domain/entity"
	"1337B04RD/internal/domain/port"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo     port.PostRepository
	commentRepo  port.CommentRepository
	imageStorage port.ImageStorage
	logger       *slog.Logger
}

// Инициализация сервиса
func NewPostService(postRepo port.PostRepository, commentRepo port.CommentRepository, imageStorage port.ImageStorage, logger *slog.Logger) *PostService {
	return &PostService{
		postRepo:     postRepo,
		commentRepo:  commentRepo,
		imageStorage: imageStorage,
		logger:       logger}
}

// // Методы для работы с постами
// (title, content string, image []byte, userID, userName, avatarURL string)
// Логика создания поста
func (s *PostService) CreatePost(ctx context.Context, userName, title, content, session_id string) (*entity.Post, error) {
	s.logger.Info("Start CreatePost",
		"title", title,
		"content", content,
	)

	newPost := entity.Post{
		PostID:       uuid.New().String(),
		UserName:     userName,
		Title:        title,
		Content:      content,
		SessionID:    session_id,
		UserAvatarID: 0, // заглушка
	}

	s.logger.Info("Creating post in repository", "postID", newPost.ID)

	post, err := s.postRepo.Create(ctx, &newPost)
	if err != nil {
		s.logger.Error("Failed to create post in repository",
			"error", err,
			"postID", newPost.ID,
		)
		return &entity.Post{}, err
	}

	s.logger.Info("Post created successfully", "postID", post.PostID)
	return post, nil
}

// func (s *PostService) UploadImage(buf []byte, filename string) (string, error) {

// }
