package storage

// internal/adapter/storage/s3.go
type S3Storage struct {
	client *minio.Client
	bucket string
}

func NewS3Storage(endpoint, accessKey, secretKey, bucket string) (*S3Storage, error) {
	// Инициализация S3 клиента
}

// Реализация методов ImageStorage
