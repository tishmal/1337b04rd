MIGRATE_VERSION=v4.16.2
MIGRATE_URL=https://github.com/golang-migrate/migrate/releases/download/$(MIGRATE_VERSION)/migrate.linux-amd64.tar.gz

MIGRATE_DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR=./migrations

include .env
export $(shell sed 's/=.*//' .env)

# Установка migrate
migrate-install:
	curl -L $(MIGRATE_URL) -o migrate.tar.gz
	tar -xzf migrate.tar.gz migrate
	chmod +x migrate
	sudo mv migrate /usr/local/bin/
	rm migrate.tar.gz
	@echo "✅ migrate установлен в /usr/local/bin"

# Применить миграции
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DB_URL)" up

# Откатить миграцию
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DB_URL)" down 1

# Удалить все таблицы
migrate-drop:
	migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DB_URL)" drop -f

# Версия
migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DB_URL)" version

run:
	go run ./cmd/1337b04rd