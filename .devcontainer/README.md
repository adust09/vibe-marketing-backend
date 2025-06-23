# GitHub Codespaces Setup

This repository is configured for GitHub Codespaces development with the following features:

## Features

- **Go 1.23** development environment
- **PostgreSQL 15** database with automatic migration setup
- **Redis 7** for caching
- **Docker-in-Docker** support
- **GitHub CLI** integration
- Pre-installed Go development tools (Air, golangci-lint, Swag, etc.)

## Services

- **API Server**: Port 8080 (automatically forwarded)
- **PostgreSQL**: Port 5432 (automatically forwarded)
- **Redis**: Port 6379 (automatically forwarded)

## Getting Started

1. Open the repository in GitHub Codespaces
2. Wait for the container to build and postCreateCommand to complete
3. The development environment will be ready with:
   - Dependencies installed via `make setup`
   - Docker services started via `make docker-compose-up`
   - Database migrations applied via `make migrate-up`

## Development Commands

Use the commands defined in the Makefile:

- `make dev` - Start development server with hot reload
- `make test` - Run tests
- `make lint` - Run linter
- `make migrate-up` - Apply database migrations

## VS Code Extensions

The following extensions are automatically installed:
- Go language support
- Docker integration
- YAML/JSON support
- GitHub Copilot
- Prettier code formatting

## Environment Variables

Copy `.env.example` to `.env` and configure your environment variables as needed. The devcontainer is pre-configured with development database and Redis connections.