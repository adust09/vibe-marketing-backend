# Ads Backend - Google Ads AI Platform API

A high-performance Go API server for the Google Ads AI Platform, built with Gin framework and PostgreSQL.

## Tech Stack

- **Framework**: Gin (Go HTTP web framework)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT with refresh tokens using golang-jwt
- **AI Integration**: Google Gemini API, Google Ads API
- **Real-time**: WebSocket support with Gorilla WebSocket
- **Testing**: Go testing package with Testify
- **Documentation**: Swagger/OpenAPI specification

## Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- Redis 7+ (for caching)
- Docker (optional, for development environment)

### Installation

1. **Clone and Setup**
   ```bash
   cd ads-backend
   go mod download
   ```

2. **Environment Configuration**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Database Setup**
   ```bash
   # Using Docker (recommended for development)
   docker-compose up -d postgres redis
   
   # Run migrations
   make migrate-up
   
   # Seed database
   make seed
   ```

4. **Start Development Server**
   ```bash
   # Install air for hot reload
   go install github.com/cosmtrek/air@latest
   
   # Start with hot reload
   make dev
   # Or without hot reload
   make run
   
   # Server runs on http://localhost:8080
   ```

## Development Commands

- `make dev` - Start with hot reload
- `make run` - Start server
- `make build` - Build binary
- `make test` - Run tests
- `make lint` - Run linter
- `make migrate-up` - Run database migrations
- `make migrate-down` - Rollback migrations

## API Documentation

Once the server is running, visit:
- Swagger UI: `http://localhost:8080/swagger/index.html`
- Health check: `http://localhost:8080/health`

## Project Structure

```
ads-backend/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── config/         # Configuration management
│   ├── controllers/    # HTTP request handlers
│   ├── services/       # Business logic layer
│   ├── models/         # Data models
│   ├── repositories/   # Data access layer
│   ├── middleware/     # HTTP middleware
│   └── routes/         # Route definitions
├── migrations/         # Database migrations
├── tests/             # Test files
└── docs/              # Documentation
```

## Contributing

1. Follow the existing code style and patterns
2. Write tests for new functionality
3. Update documentation as needed
4. Run `make lint` and `make test` before committing