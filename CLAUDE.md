# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Core Development
- `make dev` - Start development server with hot reload using Air
- `make run` - Start server without hot reload: `go run server/main.go` 
- `make build` - Build binary to `bin/server`
- `make setup` - Install development tools and copy .env.example to .env

### Testing
- `make test` - Run all tests
- `make test-verbose` - Run tests with verbose output
- `make test-coverage` - Generate test coverage report and open HTML view

### Code Quality
- `make lint` - Run golangci-lint (must run before commits)
- `make format` - Format code with go fmt

### Database Operations
- `make migrate-up` - Apply database migrations
- `make migrate-down` - Rollback database migrations  
- `make migrate-create name=migration_name` - Create new migration files
- `make seed` - Run database seeding script

### Docker
- `make docker-compose-up` - Start PostgreSQL and Redis containers
- `make docker-compose-down` - Stop containers
- `make docker-build` - Build production Docker image

## Architecture Overview

This is a Google Ads AI Platform API backend built in Go with clean architecture:

### Technology Stack
- **Framework**: Gin HTTP framework
- **Database**: PostgreSQL 14+ with GORM ORM
- **Caching**: Redis 7+
- **Authentication**: JWT with refresh tokens
- **AI Integration**: Google Gemini API, Google Ads API, Google Imagen
- **Real-time**: WebSocket support

### Project Structure
```
server/main.go              # Application entry point
api/internal/
├── config/                 # Environment configuration
├── models/                 # Data models (User, Campaign, AdGroup)
├── middleware/             # HTTP middleware (CORS, security, auth)
├── routes/                 # Route definitions
├── controllers/            # HTTP handlers (empty - needs implementation)
├── services/               # Business logic (empty - needs implementation)
├── repositories/           # Data access layer (empty - needs implementation)
├── validators/             # Input validation (empty - needs implementation)
└── utils/                  # JWT and response helpers
pkg/migrations/             # Database migrations
server/tests/               # Test organization with fixtures, mocks, helpers
```

### Data Models
- **BaseModel**: UUID-based entities with soft deletes and timestamps
- **User**: Authentication with Google Ads customer ID integration
- **Campaign**: Budget management with JSON target audience and Google Ads mapping
- **AdGroup**: Campaign subdivision for targeting
- **Analytics**: Performance metrics structure (defined but not implemented)

## Configuration

The application uses comprehensive environment-based configuration in `api/internal/config/config.go`:
- Server settings (port, environment, base URL)
- Database connection pooling and timeouts
- Redis caching configuration
- JWT token management with refresh capabilities
- Google API credentials (Ads, Gemini, Imagen)
- Email service integration
- Rate limiting and logging configuration

Copy `.env.example` to `.env` and configure before development.

## Development Notes

### Current Implementation Status
- ✅ **Complete**: Project structure, database models, migrations, authentication middleware, Docker setup
- ⚠️ **Partial**: Route structure exists but most endpoints commented out, middleware implemented but auth may need completion
- ❌ **Missing**: Controllers, services, repositories, validators, WebSocket handlers, Google API integrations, background jobs

### Key Entry Points
- `server/main.go` - Application startup and middleware setup
- `api/internal/routes/routes.go` - API routing configuration
- `api/internal/models/` - Database models and relationships
- `pkg/migrations/` - Database schema definitions

### Testing Strategy
The codebase has comprehensive test structure:
- Unit tests in `server/tests/unit/`
- Integration tests in `server/tests/integration/`
- Test fixtures in `server/tests/fixtures/`
- Mock implementations in `server/tests/mocks/`
- Test helpers in `server/tests/helpers/`

Always run `make lint` and `make test` before committing changes.

### Development Workflow
1. Start dependencies: `make docker-compose-up`
2. Run migrations: `make migrate-up`
3. Start development server: `make dev`
4. Server runs on http://localhost:8080
5. Health check: http://localhost:8080/health
6. API base path: /api/v1

### Security Features
- JWT authentication with refresh tokens
- CORS middleware configured
- Security headers middleware
- Password hashing structure in place
- Rate limiting configuration available
- Environment-based secrets management

## Import Paths

The Go module name is `ads-backend`. When importing internal packages, use:
```go
import "ads-backend/internal/config"
import "ads-backend/internal/models"
```

The correct server entry point is `server/main.go`, not `cmd/server/main.go` as mentioned in some documentation.