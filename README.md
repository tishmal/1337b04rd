# 1337b04rd 📦

> **Anonymous image board backend written in Go**

---

## ✨ Features

- 🛸 Anonymous sessions with Rick & Morty avatars
- 📍 Thread and comment posting
- 📷 Image uploads to MinIO (S3-compatible)
- ⌛ Auto-cleanup of threads without comments
- ⚖️ Moderation-ready architecture
- ✨ Clean, layered Go codebase (Hexagonal Architecture)

---

## 📚 Tech Stack

- **Go** 1.23+
- **PostgreSQL**
- **MinIO** (S3-compatible object storage)
- **HTML Templates** (server-rendered)
- **TailwindCSS** (frontend styling)
- **Docker** (local development)

---

## 🧬 Project Structure

```
1337b04rd/
├── cmd/               # CLI entrypoint (main.go)
├── config/            # Manual .env parsing & configuration
├── db/                # SQL init scripts
├── internal/          # Core business logic and adapters
│   ├── adapters/      # Infrastructure adapters (db, http, s3)
│   │   ├── http/      # HTTP handlers & middleware
│   │   ├── postgres/  # PostgreSQL repository implementation
│   │   ├── rickmorty/ # External Rick & Morty avatar client
│   │   └── s3/        # MinIO/S3 storage client
│   ├── app/
│   │   ├── common/    # Utilities (UUIDs, logger)
│   │   ├── ports/     # Interfaces for dependency injection
│   │   └── services/  # Business logic implementation
│   └── domain/        # Domain models & rules
│       ├── avatar/
│       ├── comment/
│       ├── errors/
│       ├── session/
│       └── thread/
├── test/              # Unit & integration tests
│   ├── integration/
│   ├── testdata/
│   └── unit/
└── web/               # Frontend (templates + static assets)
    ├── static/        # Static files (images, CSS)
    └── templates/     # HTML templates
```

---

## 🚀 Getting Started

```bash
# Clone the repository
git clone https://your-repo-url/1337b04rd.git

# Move into project directory
cd 1337b04rd

# Build and run
docker-compose up --build
```

---

## 🛠️ Environment Variables

Create a `.env` file:

```bash
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=1337b04rd
DB_SSLMODE=disable

S3_ENDPOINT=localhost:9000
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_BUCKET_THREADS=threads
S3_BUCKET_COMMENTS=comments
S3_REGION=us-east-1
S3_USE_SSL=false

SESSION_COOKIE_NAME=1337session
SESSION_DURATION_DAYS=7

AVATAR_API_BASE_URL=https://rickandmortyapi.com/api/character
APP_ENV=development
```

---

## 🧹 Code Quality

- Hexagonal Architecture
- Clean, layered design
- Structured logging (`slog`)
- Environment-based configuration

---

## 📜 License

Distributed under the MIT License.
See `LICENSE` for more information.