package port

type ImageStorage interface {
	Upload(data []byte, filename string) (string, error)
	Get(filename string) ([]byte, error)
	Delete(filename string) error
}
