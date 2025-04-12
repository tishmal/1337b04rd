package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client        *minio.Client
	postBucket    string
	commentBucket string
	logger        *slog.Logger
}

func NewS3Storage(endpoint, accessKey, secretKey string, useSSL bool, postBucket, commentBucket string, logger *slog.Logger) (*S3Storage, error) {
	// Инициализация клиента MinIO
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client: %w", err)
	}

	s3 := &S3Storage{
		client:        client,
		postBucket:    postBucket,
		commentBucket: commentBucket,
		logger:        logger,
	}

	// Проверка существования бакетов и их создание, если они не существуют
	ctx := context.Background()

	// Проверка и создание бакета для постов
	exists, err := client.BucketExists(ctx, postBucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check post bucket existence: %w", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, postBucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create post bucket: %w", err)
		}
		logger.Info("Created post bucket", "bucket", postBucket)
	}

	// Проверка и создание бакета для комментариев
	exists, err = client.BucketExists(ctx, commentBucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check comment bucket existence: %w", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, commentBucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create comment bucket: %w", err)
		}
		logger.Info("Created comment bucket", "bucket", commentBucket)
	}

	return s3, nil
}

// Реализация методов ImageStorage
