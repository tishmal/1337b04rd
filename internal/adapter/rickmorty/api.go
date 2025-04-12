package rickmorty

// internal/adapter/rickmorty/api.go
type RickMortyAPI struct {
	baseURL string
	cache   map[int]Character // Кэш персонажей
}

func NewRickMortyAPI(baseURL string) *RickMortyAPI {
	// Инициализация API клиента
	return &RickMortyAPI{}
	// заглушка
}

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Добавьте другие поля, которые вам нужны для персонажа
}

// Реализация методов AvatarService
