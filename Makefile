.PHONY: up down logs db-up db-reset test test-reset build fmt sqlc

COMPOSE := docker compose
DB_SERVICE := database

up:
	$(COMPOSE) up --build

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f

db-up:
	$(COMPOSE) up -d $(DB_SERVICE)

db-reset:
	$(COMPOSE) down -v
	$(COMPOSE) up -d $(DB_SERVICE)

test: db-up
	@until $(COMPOSE) exec -T $(DB_SERVICE) pg_isready -U user -d mydatabase >/dev/null 2>&1; do \
		echo "Esperando a PostgreSQL..."; \
		sleep 1; \
	done
	go test ./... -v -count=1

test-reset: db-reset
	@$(MAKE) test

build:
	go build ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

sqlc:
	sqlc generate