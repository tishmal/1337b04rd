1337b04rd/
├── cmd/
│   └── 1337b04rd/
│       └── main.go           # Точка входа в приложение
├── internal/
│   ├── domain/               # Домен (ядро)
│   │   ├── entity/           # Сущности
│   │   ├── port/             # Интерфейсы (порты)
│   │   └── service/          # Сервисы (бизнес-логика)
│   ├── adapter/              # Адаптеры
│   │   ├── db/               # Адаптер PostgreSQL
│   │   ├── storage/          # Адаптер S3
│   │   ├── rickmorty/        # Адаптер Rick and Morty API
│   │   └── http/             # HTTP-адаптер
│   └── config/               # Конфигурация
├── web/                      # Шаблоны и статические файлы
│   ├── templates/            # HTML-шаблоны
│   └── static/               # CSS, JS, изображения
├── test/                     # Тесты
└── go.mod, go.sum            # Зависимости