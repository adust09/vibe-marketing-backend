.PHONY: dev run build build-linux test test-verbose test-coverage migrate-up migrate-down migrate-create seed lint format docker-build docker-run generate mock swagger clean

# Development
dev:
	air

run:
	go run cmd/server/main.go

# Build
build:
	go build -o bin/server cmd/server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/server-linux cmd/server/main.go

# Testing
test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Database
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

seed:
	go run scripts/seed.go

# Code Quality
lint:
	golangci-lint run

format:
	go fmt ./...

# Docker
docker-build:
	docker build -t ads-backend .

docker-run:
	docker run -p 8080:8080 ads-backend

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

# Generate
generate:
	go generate ./...

mock:
	mockgen -source=internal/repositories/interfaces.go -destination=tests/mocks/repositories.go

swagger:
	swag init -g cmd/server/main.go -o api/

# Dependencies
deps:
	go mod download
	go mod tidy

# Clean
clean:
	rm -rf bin/
	rm -rf coverage.out
	rm -rf api/docs.go api/swagger.json api/swagger.yaml

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golang/mock/mockgen@latest

# Setup development environment
setup:
	make install-tools
	cp .env.example .env
	go mod download