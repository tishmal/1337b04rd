1337b04rd 📂

🧬 Project Structure

📚 Tech Stack

    Go 1.23+
    PostgreSQL
    MinIO (S3-compatible object storage)
    HTML templates (server-rendered)
    TailwindCSS (frontend styling)
    Docker (virtualization)

🧬 Project Structure

1337b04rd/
├── cmd/                     # CLI entrypoint
├── config/                  # Manual .env parsing & configuration
├── db/                      # SQL init scripts
├── internal/                # Core business logic and adapters
│   ├── adapters/            # Infrastructure adapters
│   │   ├── http/            # HTTP handlers & middleware
│   │   ├── postgres/        # PostgreSQL repositories
│   │   ├── rickmorty/       # External avatar client
│   │   └── s3/              # S3/MinIO storage client
│   ├── app/
│   │   ├── common/
│   │   │   ├── logger/      # Structured logging
│   │   │   └── utils/       # Helpers (UUID, etc)
│   │   ├── ports/           # Interfaces for domain-driven services
│   │   └── services/        # Business logic
│   └── domain/              # Core domain models and rules
│       ├── avatar/
│       ├── comment/
│       ├── errors/
│       ├── session/
│       └── thread/
├── test/                    # Tests and testdata
│   ├── integration/
│   ├── testdata/
│   └── unit/
└── web/                     # Frontend
    ├── static/              # Static assets (img)
    │   └── img/
    └── templates/           # HTML templates