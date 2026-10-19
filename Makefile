.PHONY: up generate down logs db-up db-reset test test-api test-reset build fmt

COMPOSE := docker compose
DB_SERVICE := database
SQLC_IMAGE := sqlc/sqlc:1.31.1

up: generate
	$(COMPOSE) up -d --build

generate:
	docker run --rm -v "$$(pwd):/src" -w /src $(SQLC_IMAGE) generate

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f

db-up:
	$(COMPOSE) up -d $(DB_SERVICE)

db-reset:
	$(COMPOSE) down -v
	$(COMPOSE) up -d $(DB_SERVICE)

test: test-api

test-api: generate build
	@set -eu; \
		trap '$(COMPOSE) down -v --remove-orphans' EXIT INT TERM; \
		$(COMPOSE) down -v --remove-orphans; \
		$(COMPOSE) up -d --build; \
		echo "Esperando a PostgreSQL..."; \
		until $(COMPOSE) exec -T $(DB_SERVICE) pg_isready -U user -d mydatabase >/dev/null 2>&1; do \
			sleep 1; \
		done; \
		echo "Esperando a la API..."; \
		until curl -fsS http://localhost:8080/ >/dev/null 2>&1; do \
			sleep 1; \
		done; \
		./requests.sh

test-reset: db-reset
	@$(MAKE) test

build:
	go build ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')