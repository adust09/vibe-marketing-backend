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
server/main.go              # Application entry point with database initialization
api/internal/
├── config/                 # Environment configuration
├── database/               # Database connection and setup
├── models/                 # Data models (User, Campaign, AdGroup, Keyword)
├── middleware/             # HTTP middleware (CORS, security, auth)
├── routes/                 # Route definitions with CPC endpoints
├── controllers/            # HTTP handlers for CPC operations
├── services/               # Business logic with Google Ads integration
├── repositories/           # Data access layer with CRUD operations
├── validators/             # Input validation (empty - needs implementation)
└── utils/                  # JWT and response helpers
pkg/migrations/             # Database migrations including CPC fields
server/tests/               # Test organization with fixtures, mocks, helpers
```

### Data Models
- **BaseModel**: UUID-based entities with soft deletes and timestamps
- **User**: Authentication with Google Ads customer ID integration
- **Campaign**: Budget management with JSON target audience, Google Ads mapping, and CPC tracking (cpc, average_cpc, max_cpc)
- **AdGroup**: Campaign subdivision for targeting with CPC metrics
- **Keyword**: Keyword-level tracking with CPC, quality score, impressions, clicks, and cost data
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
- ✅ **Complete**: Project structure, database models with CPC fields, migrations, authentication middleware, Docker setup, CPC API endpoints
- ✅ **Complete**: Controllers, services, repositories for Campaign/AdGroup/Keyword management
- ✅ **Complete**: Google Ads API service foundation with mock data
- ⚠️ **Partial**: Google Ads API actual integration (foundation ready), middleware implemented but auth may need completion
- ❌ **Missing**: Validators, WebSocket handlers, full Google API integrations, background jobs

### Key Entry Points
- `server/main.go` - Application startup with database initialization and middleware setup
- `api/internal/routes/routes.go` - API routing configuration with CPC endpoints
- `api/internal/models/` - Database models and relationships including CPC fields
- `api/internal/controllers/` - HTTP handlers for Campaign, AdGroup, and Keyword operations
- `api/internal/services/` - Business logic with Google Ads API integration
- `api/internal/repositories/` - Data access layer with CRUD operations
- `pkg/migrations/` - Database schema definitions including CPC tables

### Testing Strategy
The codebase has comprehensive test structure:
- Unit tests in `server/tests/unit/` with CPC service tests implemented
- Integration tests in `server/tests/integration/`
- Test fixtures in `server/tests/fixtures/`
- Mock implementations in `server/tests/mocks/`
- Test helpers in `server/tests/helpers/`
- Mock repositories and Google Ads service for isolated testing

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

## CPC (Cost Per Click) API Endpoints

### Campaign CPC Management
- `GET /api/v1/campaigns` - Get all campaigns with CPC data
- `GET /api/v1/campaigns/:id` - Get specific campaign with CPC metrics
- `PUT /api/v1/campaigns/:id/cpc` - Update campaign CPC values manually
- `POST /api/v1/campaigns/:id/refresh-cpc` - Refresh CPC data from Google Ads API

### AdGroup CPC Management
- `GET /api/v1/campaigns/:campaign_id/adgroups` - Get ad groups for a campaign
- `GET /api/v1/adgroups/:id` - Get specific ad group with CPC metrics
- `PUT /api/v1/adgroups/:id/cpc` - Update ad group CPC values manually
- `POST /api/v1/adgroups/:id/refresh-cpc` - Refresh CPC data from Google Ads API

### Keyword CPC Management
- `GET /api/v1/adgroups/:adgroup_id/keywords` - Get keywords for an ad group
- `GET /api/v1/keywords/:id` - Get specific keyword with full metrics (CPC, quality score, impressions, clicks, cost)
- `PUT /api/v1/keywords/:id/cpc` - Update keyword CPC values manually
- `POST /api/v1/keywords/:id/refresh-cpc` - Refresh keyword data from Google Ads API

### CPC Data Structure
All CPC endpoints return the following fields:
- `cpc` - Current cost per click (decimal)
- `average_cpc` - Average cost per click over time period (decimal)
- `max_cpc` - Maximum bid for cost per click (decimal)

Keywords additionally include:
- `quality_score` - Google Ads quality score (1-10)
- `impressions` - Total impressions count
- `clicks` - Total clicks count
- `cost` - Total cost spent

### Google Ads Integration
The system includes a `GoogleAdsService` that handles:
- Authentication with Google Ads API
- Fetching real-time CPC data
- Updating local database with fresh metrics
- Error handling and retry logic

Current implementation uses mock data for development. To enable real Google Ads integration:
1. Configure `GOOGLE_ADS_DEVELOPER_TOKEN` in environment
2. Set up OAuth2 credentials in `GOOGLE_CLIENT_ID` and `GOOGLE_CLIENT_SECRET`
3. Implement actual API calls in `api/internal/services/google_ads_service.go`