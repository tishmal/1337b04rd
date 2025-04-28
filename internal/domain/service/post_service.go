package service

import (
	"1337B04RD/internal/app/common/logger"
	"1337B04RD/internal/app/common/utils"
	"1337B04RD/internal/domain/entity"
	"1337B04RD/internal/domain/port"
	"fmt"

	"context"
)

type PostService struct {
	postRepo    port.PostRepository
	commentRepo port.CommentRepository
}

// Инициализация сервиса
func NewPostService(postRepo port.PostRepository, commentRepo port.CommentRepository) *PostService {
	return &PostService{
		postRepo:    postRepo,
		commentRepo: commentRepo}
}

// // Методы для работы с постами
// (title, content string, image []byte, userID, userName, avatarURL string)
// Логика создания поста
func (s *PostService) CreatePost(ctx context.Context, userName, title, content, session_id string) (*entity.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("err Empty Title")
	}

	if content == "" {
		return nil, fmt.Errorf("err Empty contnent")
	}

	logger.Info("Start CreatePost",
		"title", title,
		"content", content,
	)

	id, err := utils.NewUUID()
	if err != nil {

	}

	newPost := entity.Post{
		PostID:       id,
		UserName:     userName,
		Title:        title,
		Content:      content,
		SessionID:    session_id,
		UserAvatarID: 0, // заглушка
	}

	logger.Info("Creating post in repository", "postID", newPost.ID)

	post, err := s.postRepo.Create(ctx, &newPost)
	if err != nil {
		logger.Error("Failed to create post in repository",
			"error", err,
			"postID", newPost.ID,
		)
		return &entity.Post{}, err
	}

	logger.Info("Post created successfully", "postID", post.PostID)
	return post, nil
}

// func (s *PostService) UploadImage(buf []byte, filename string) (string, error) {

// }
