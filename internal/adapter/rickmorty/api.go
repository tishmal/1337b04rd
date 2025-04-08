package rickmorty

// internal/adapter/rickmorty/api.go
type RickMortyAPI struct {
	baseURL string
	cache   map[int]Character // Кэш персонажей
}

func NewRickMortyAPI(baseURL string) *RickMortyAPI {
	// Инициализация API клиента
}

// Реализация методов AvatarService
