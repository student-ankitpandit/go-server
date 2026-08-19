.PHONY: build run build-migrate migrate-up migrate-down

build:
	@go build -o bin/api ./cmd/api

build-migrate: 
	@go build -o bin/migrate ./cmd/migrate
	
run: build
	@./bin/api

migrate-up: build-migrate
	@go run ./cmd/migrate up
migrate-down:
	@go run ./cmd/migrate down