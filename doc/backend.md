# Backend Repository - Google Ads AI Platform

This document outlines the plan and implementation details for separating the backend from the current Next.js full-stack application into a dedicated Go API server.

## Table of Contents

1. [Overview](#overview)
2. [Repository Structure](#repository-structure)
3. [Migration Plan](#migration-plan)
4. [Development Setup](#development-setup)
5. [API Endpoints](#api-endpoints)
6. [Database Design](#database-design)
7. [Authentication & Authorization](#authentication--authorization)
8. [AI Services Integration](#ai-services-integration)
9. [Testing Strategy](#testing-strategy)
10. [Deployment & DevOps](#deployment--devops)
11. [Environment Configuration](#environment-configuration)

## Overview

The backend will be a RESTful API server built with:
- **Framework**: Gin (Go HTTP web framework)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT with refresh tokens using golang-jwt
- **AI Integration**: Google Gemini API, Google Ads API
- **Real-time**: WebSocket support with Gorilla WebSocket
- **Testing**: Go testing package with Testify
- **Documentation**: Swagger/OpenAPI specification with gin-swagger

### Key Features
- Campaign management CRUD operations
- AI-powered ad generation and optimization
- Real-time analytics and performance tracking
- User authentication and authorization
- Google Ads API integration
- WebSocket support for live updates
- High-performance concurrent request handling

## Repository Structure

```
ads-backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration management
│   │   ├── config.go
│   │   └── database.go
│   ├── controllers/             # HTTP request handlers
│   │   ├── auth_controller.go
│   │   ├── campaign_controller.go
│   │   ├── analytics_controller.go
│   │   ├── ai_controller.go
│   │   └── user_controller.go
│   ├── services/                # Business logic layer
│   │   ├── auth_service.go
│   │   ├── campaign_service.go
│   │   ├── analytics_service.go
│   │   ├── gemini_service.go
│   │   ├── google_ads_service.go
│   │   ├── imagen_service.go
│   │   └── user_service.go
│   ├── models/                  # Data models and structs
│   │   ├── user.go
│   │   ├── campaign.go
│   │   ├── adgroup.go
│   │   ├── analytics.go
│   │   └── base.go
│   ├── repositories/            # Data access layer
│   │   ├── user_repository.go
│   │   ├── campaign_repository.go
│   │   ├── analytics_repository.go
│   │   └── interfaces.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth_middleware.go
│   │   ├── cors_middleware.go
│   │   ├── error_middleware.go
│   │   ├── validation_middleware.go
│   │   └── rate_limit_middleware.go
│   ├── routes/                  # Route definitions
│   │   ├── auth_routes.go
│   │   ├── campaign_routes.go
│   │   ├── analytics_routes.go
│   │   ├── ai_routes.go
│   │   ├── user_routes.go
│   │   └── routes.go
│   ├── validators/              # Request validation
│   │   ├── auth_validator.go
│   │   ├── campaign_validator.go
│   │   └── common_validator.go
│   ├── utils/                   # Shared utilities
│   │   ├── logger.go
│   │   ├── response.go
│   │   ├── crypto.go
│   │   ├── jwt.go
│   │   └── constants.go
│   ├── websocket/               # WebSocket handlers
│   │   ├── handlers/
│   │   │   ├── analytics_handler.go
│   │   │   └── campaign_handler.go
│   │   ├── middleware/
│   │   │   └── ws_auth_middleware.go
│   │   └── hub.go
│   └── jobs/                    # Background jobs
│       ├── analytics_sync.go
│       └── campaign_optimization.go
├── pkg/                         # Public packages
│   ├── httputil/
│   │   └── response.go
│   └── timeutil/
│       └── time.go
├── api/                         # API documentation
│   ├── swagger.yaml
│   └── docs.go                  # Generated Swagger docs
├── scripts/                     # Build and deployment scripts
│   ├── build.sh
│   ├── deploy.sh
│   ├── migrate.sh
│   └── seed.go
├── migrations/                  # Database migrations
│   ├── 001_create_users_table.up.sql
│   ├── 001_create_users_table.down.sql
│   ├── 002_create_campaigns_table.up.sql
│   ├── 002_create_campaigns_table.down.sql
│   └── ...
├── tests/                       # Test files
│   ├── unit/
│   │   ├── services/
│   │   ├── controllers/
│   │   └── utils/
│   ├── integration/
│   │   ├── api/
│   │   └── database/
│   ├── mocks/                   # Generated mocks
│   ├── fixtures/                # Test data
│   └── helpers/
├── docs/                        # Documentation
│   ├── api/
│   ├── database/
│   └── deployment/
├── .env.example                 # Environment variables template
├── .gitignore
├── .golangci.yml               # Linter configuration
├── docker-compose.yml          # Development environment
├── Dockerfile                  # Production container
├── go.mod                      # Go modules
├── go.sum                      # Go modules checksum
├── Makefile                    # Build commands
└── README.md
```

## Migration Plan

### Phase 1: Initial Setup
1. **Create New Go Repository**
   - Initialize Go module: `go mod init ads-backend`
   - Setup project structure following Go standards
   - Configure development tools (golangci-lint, air for hot reload)
   - Setup testing framework with testify

2. **Basic Gin Server**
   - Create main.go with Gin HTTP server
   - Setup middleware (CORS, recovery, logger)
   - Configure environment variables with viper
   - Add health check endpoint

### Phase 2: Database Layer
1. **PostgreSQL Setup**
   - Configure GORM with PostgreSQL driver
   - Create database models (User, Campaign, AdGroup, Analytics)
   - Setup database migrations with golang-migrate
   - Implement connection pooling

2. **Repository Pattern**
   - Create repository interfaces
   - Implement repository structs with GORM
   - Add database transaction support
   - Setup database testing with testcontainers

### Phase 3: Core API Migration
1. **Move API Routes**
   - Convert Next.js API handlers to Gin controllers
   - Implement request validation with go-playground/validator
   - Add proper error handling with custom error types
   - Setup response formatting

2. **Business Logic Layer**
   - Create service layer for business logic
   - Implement dependency injection
   - Add logging with structured logging (logrus/zap)
   - Setup configuration management

### Phase 4: AI Services Migration
1. **Gemini API Service**
   - Port TypeScript Gemini client to Go
   - Implement HTTP client with retry logic
   - Add response caching with Redis/go-redis
   - Setup rate limiting with go-rate

2. **Google Ads API Service**
   - Implement Google Ads API client in Go
   - Setup OAuth2 flow with golang.org/x/oauth2
   - Add API response caching
   - Implement error handling and retries

### Phase 5: Authentication & Security
1. **JWT Authentication**
   - Implement JWT with golang-jwt/jwt
   - Add refresh token mechanism
   - Setup password hashing with bcrypt
   - Implement role-based access control

2. **Security Middleware**
   - Configure CORS with gin-cors
   - Add rate limiting with gin-rate-limit
   - Implement input sanitization
   - Setup security headers

### Phase 6: Real-time Features & Testing
1. **WebSocket Server**
   - Setup WebSocket with gorilla/websocket
   - Implement real-time analytics updates
   - Add campaign status notifications
   - Create WebSocket hub for connection management

2. **Comprehensive Testing**
   - Unit tests for all services and controllers
   - Integration tests with testcontainers
   - API testing with httptest
   - Mock generation with gomock

## Development Setup

### Prerequisites
- Go 1.21+ 
- PostgreSQL 14+
- Redis 7+ (for caching)
- Docker (optional, for development environment)
- Make (for build commands)

### Installation Steps

1. **Clone and Setup**
   ```bash
   git clone <backend-repo-url>
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

### Development Commands (Makefile)
```makefile
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

# Generate
generate:
	go generate ./...

mock:
	mockgen -source=internal/repositories/interfaces.go -destination=tests/mocks/repositories.go

swagger:
	swag init -g cmd/server/main.go -o api/
```

## API Endpoints

### Go Gin Route Examples

```go
// internal/routes/auth_routes.go
func SetupAuthRoutes(r *gin.RouterGroup, authController *controllers.AuthController) {
    auth := r.Group("/auth")
    {
        auth.POST("/register", authController.Register)
        auth.POST("/login", authController.Login)
        auth.POST("/refresh", authController.RefreshToken)
        auth.POST("/logout", middleware.AuthMiddleware(), authController.Logout)
        auth.POST("/forgot-password", authController.ForgotPassword)
        auth.POST("/reset-password", authController.ResetPassword)
        auth.GET("/profile", middleware.AuthMiddleware(), authController.GetProfile)
        auth.PUT("/profile", middleware.AuthMiddleware(), authController.UpdateProfile)
    }
}

// internal/routes/campaigns_routes.go
func SetupCampaignRoutes(r *gin.RouterGroup, campaignController *controllers.CampaignController) {
    campaigns := r.Group("/campaigns")
    campaigns.Use(middleware.AuthMiddleware())
    {
        campaigns.GET("", campaignController.ListCampaigns)
        campaigns.POST("", campaignController.CreateCampaign)
        campaigns.GET("/:id", campaignController.GetCampaign)
        campaigns.PUT("/:id", campaignController.UpdateCampaign)
        campaigns.DELETE("/:id", campaignController.DeleteCampaign)
        campaigns.POST("/:id/start", campaignController.StartCampaign)
        campaigns.POST("/:id/pause", campaignController.PauseCampaign)
        campaigns.POST("/:id/stop", campaignController.StopCampaign)
        
        // Ad Groups
        campaigns.GET("/:id/ad-groups", campaignController.ListAdGroups)
        campaigns.POST("/:id/ad-groups", campaignController.CreateAdGroup)
    }
}
```

### Complete API Endpoints

#### Authentication
```
POST   /api/auth/register       # User registration
POST   /api/auth/login          # User login
POST   /api/auth/refresh        # Refresh JWT token
POST   /api/auth/logout         # User logout
POST   /api/auth/forgot-password # Password reset request
POST   /api/auth/reset-password  # Password reset
GET    /api/auth/profile        # Get user profile
PUT    /api/auth/profile        # Update user profile
```

#### Campaigns
```
GET    /api/campaigns           # List campaigns (with filtering)
POST   /api/campaigns           # Create new campaign
GET    /api/campaigns/:id       # Get campaign details
PUT    /api/campaigns/:id       # Update campaign
DELETE /api/campaigns/:id       # Delete campaign
POST   /api/campaigns/:id/start # Start campaign
POST   /api/campaigns/:id/pause # Pause campaign
POST   /api/campaigns/:id/stop  # Stop campaign
```

#### AI Services
```
POST   /api/ai/generate-ad-text            # Generate ad headlines/descriptions
POST   /api/ai/generate-keywords           # Generate keyword suggestions
POST   /api/ai/generate-response-parts     # Generate response part ads
POST   /api/ai/analyze-performance         # Analyze campaign performance
POST   /api/ai/optimization-suggestions    # Get optimization recommendations
POST   /api/ai/generate-images             # Generate ad images (Imagen API)
```

#### Analytics
```
GET    /api/analytics/overview             # Dashboard overview
GET    /api/analytics/campaigns/:id        # Campaign analytics
GET    /api/analytics/performance          # Performance reports
GET    /api/analytics/real-time            # Real-time metrics
POST   /api/analytics/reports              # Generate custom reports
```

## Database Design

### GORM Models

```go
// internal/models/user.go
type User struct {
    ID                    uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Email                 string     `gorm:"uniqueIndex;not null" json:"email"`
    PasswordHash          string     `gorm:"not null" json:"-"`
    FirstName             *string    `json:"first_name"`
    LastName              *string    `json:"last_name"`
    Role                  string     `gorm:"default:user" json:"role"`
    IsActive              bool       `gorm:"default:true" json:"is_active"`
    EmailVerified         bool       `gorm:"default:false" json:"email_verified"`
    GoogleAdsCustomerID   *string    `json:"google_ads_customer_id"`
    Campaigns             []Campaign `gorm:"foreignKey:UserID" json:"campaigns,omitempty"`
    CreatedAt             time.Time  `json:"created_at"`
    UpdatedAt             time.Time  `json:"updated_at"`
}

// internal/models/campaign.go
type Campaign struct {
    ID                    uuid.UUID           `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    UserID                uuid.UUID           `gorm:"type:uuid;not null" json:"user_id"`
    Name                  string              `gorm:"not null" json:"name"`
    Status                string              `gorm:"default:draft" json:"status"`
    BudgetAmount          decimal.Decimal     `gorm:"type:decimal(12,2);not null" json:"budget_amount"`
    DailyBudget           decimal.Decimal     `gorm:"type:decimal(12,2);not null" json:"daily_budget"`
    StartDate             *time.Time          `gorm:"type:date" json:"start_date"`
    EndDate               *time.Time          `gorm:"type:date" json:"end_date"`
    Objective             string              `gorm:"not null" json:"objective"`
    TargetAudience        datatypes.JSON      `json:"target_audience"`
    GoogleAdsCampaignID   *string             `json:"google_ads_campaign_id"`
    Platform              string              `gorm:"default:google" json:"platform"`
    User                  User                `gorm:"foreignKey:UserID" json:"user,omitempty"`
    AdGroups              []AdGroup           `gorm:"foreignKey:CampaignID" json:"ad_groups,omitempty"`
    Analytics             []CampaignAnalytics `gorm:"foreignKey:CampaignID" json:"analytics,omitempty"`
    CreatedAt             time.Time           `json:"created_at"`
    UpdatedAt             time.Time           `json:"updated_at"`
}

// internal/models/adgroup.go
type AdGroup struct {
    ID                  uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    CampaignID          uuid.UUID      `gorm:"type:uuid;not null" json:"campaign_id"`
    Name                string         `gorm:"not null" json:"name"`
    Status              string         `gorm:"default:active" json:"status"`
    Targeting           datatypes.JSON `json:"targeting"`
    GoogleAdsAdGroupID  *string        `json:"google_ads_ad_group_id"`
    Campaign            Campaign       `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
    CreatedAt           time.Time      `json:"created_at"`
    UpdatedAt           time.Time      `json:"updated_at"`
}

// internal/models/analytics.go
type CampaignAnalytics struct {
    ID           uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    CampaignID   uuid.UUID       `gorm:"type:uuid;not null" json:"campaign_id"`
    Date         time.Time       `gorm:"type:date;not null" json:"date"`
    Impressions  int             `gorm:"default:0" json:"impressions"`
    Clicks       int             `gorm:"default:0" json:"clicks"`
    Conversions  int             `gorm:"default:0" json:"conversions"`
    CostMicros   int64           `gorm:"default:0" json:"cost_micros"`
    RevenueMicros *int64         `json:"revenue_micros"`
    Campaign     Campaign        `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
    CreatedAt    time.Time       `json:"created_at"`
}
```

### Database Migrations

```sql
-- migrations/001_create_users_table.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    email_verified BOOLEAN DEFAULT false,
    google_ads_customer_id VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrations/002_create_campaigns_table.up.sql
CREATE TABLE campaigns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'draft',
    budget_amount DECIMAL(12,2) NOT NULL,
    daily_budget DECIMAL(12,2) NOT NULL,
    start_date DATE,
    end_date DATE,
    objective VARCHAR(100) NOT NULL,
    target_audience JSONB,
    google_ads_campaign_id VARCHAR(50),
    platform VARCHAR(50) DEFAULT 'google',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Authentication & Authorization

### JWT Implementation with Go

```go
// internal/utils/jwt.go
package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func GenerateTokens(userID, email, role string) (accessToken, refreshToken string, err error) {
    // Access Token (15 minutes)
    accessClaims := &Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessToken, err = accessTokenObj.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", "", err
    }

    // Refresh Token (7 days)
    refreshClaims := &Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshToken, err = refreshTokenObj.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
    
    return accessToken, refreshToken, err
}

func ValidateToken(tokenString, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    
    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, jwt.ErrInvalidKey
}
```

### Authentication Middleware

```go
// internal/middleware/auth_middleware.go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        claims, err := utils.ValidateToken(tokenString, os.Getenv("JWT_SECRET"))
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Set user info in context
        c.Set("user_id", claims.UserID)
        c.Set("user_email", claims.Email)
        c.Set("user_role", claims.Role)
        c.Next()
    }
}
```

## AI Services Integration

### Gemini API Service in Go

```go
// internal/services/gemini_service.go
package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

type GeminiService struct {
    client  *http.Client
    apiKey  string
    baseURL string
}

type GeminiRequest struct {
    Model    string `json:"model"`
    Contents string `json:"contents"`
    Config   struct {
        Temperature      float64 `json:"temperature"`
        MaxOutputTokens  int     `json:"maxOutputTokens"`
    } `json:"config"`
}

type GeminiResponse struct {
    Text string `json:"text"`
}

func NewGeminiService() *GeminiService {
    return &GeminiService{
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
        apiKey:  os.Getenv("GEMINI_API_KEY"),
        baseURL: "https://generativelanguage.googleapis.com/v1",
    }
}

func (s *GeminiService) GenerateAdText(productInfo, targetAudience, adType string) (*AdTextResponse, error) {
    prompt := s.buildAdTextPrompt(productInfo, targetAudience, adType)
    
    request := GeminiRequest{
        Model:    "gemini-2.0-flash-live-001",
        Contents: prompt,
        Config: struct {
            Temperature     float64 `json:"temperature"`
            MaxOutputTokens int     `json:"maxOutputTokens"`
        }{
            Temperature:     0.8,
            MaxOutputTokens: 1024,
        },
    }

    jsonData, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal request: %w", err)
    }

    url := fmt.Sprintf("%s/models/generateContent?key=%s", s.baseURL, s.apiKey)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := s.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to make request: %w", err)
    }
    defer resp.Body.Close()

    var geminiResp GeminiResponse
    if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return s.parseAdTextResponse(geminiResp.Text, adType), nil
}

type AdTextResponse struct {
    Headlines     []string `json:"headlines,omitempty"`
    Descriptions  []string `json:"descriptions,omitempty"`
    ResponseParts []string `json:"response_parts,omitempty"`
}

func (s *GeminiService) parseAdTextResponse(text, adType string) *AdTextResponse {
    // Implementation similar to TypeScript version but in Go
    // Parse the AI response and extract headlines, descriptions, response parts
    response := &AdTextResponse{}
    
    // Parse logic here...
    
    return response
}
```

### Google Ads API Service in Go

```go
// internal/services/google_ads_service.go
package services

import (
    "context"
    "fmt"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

type GoogleAdsService struct {
    config       *oauth2.Config
    client       *http.Client
    developerToken string
}

func NewGoogleAdsService() *GoogleAdsService {
    config := &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        Scopes:       []string{"https://www.googleapis.com/auth/adwords"},
        Endpoint:     google.Endpoint,
    }

    return &GoogleAdsService{
        config:         config,
        developerToken: os.Getenv("GOOGLE_ADS_DEVELOPER_TOKEN"),
    }
}

func (s *GoogleAdsService) GetCampaignPerformance(ctx context.Context, customerID, campaignID string, token *oauth2.Token) (*CampaignPerformance, error) {
    client := s.config.Client(ctx, token)
    
    query := fmt.Sprintf(`
        SELECT
            campaign.id,
            campaign.name,
            campaign.status,
            metrics.impressions,
            metrics.clicks,
            metrics.cost_micros,
            metrics.conversions
        FROM campaign
        WHERE campaign.id = %s
    `, campaignID)

    // Make API request
    url := fmt.Sprintf("https://googleads.googleapis.com/v14/customers/%s/googleAds:search", customerID)
    
    // Implementation continues...
    
    return nil, nil
}

type CampaignPerformance struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Status      string `json:"status"`
    Impressions int64  `json:"impressions"`
    Clicks      int64  `json:"clicks"`
    CostMicros  int64  `json:"cost_micros"`
    Conversions int64  `json:"conversions"`
}
```

## Testing Strategy

### Unit Testing with Testify

```go
// tests/unit/services/campaign_service_test.go
package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "ads-backend/internal/services"
    "ads-backend/tests/mocks"
)

func TestCampaignService_CreateCampaign(t *testing.T) {
    // Arrange
    mockRepo := new(mocks.MockCampaignRepository)
    service := services.NewCampaignService(mockRepo)
    
    campaign := &models.Campaign{
        Name:         "Test Campaign",
        BudgetAmount: decimal.NewFromFloat(1000.00),
        DailyBudget:  decimal.NewFromFloat(50.00),
        Objective:    "conversions",
    }
    
    mockRepo.On("Create", mock.AnythingOfType("*models.Campaign")).Return(nil)
    
    // Act
    result, err := service.CreateCampaign(campaign)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "Test Campaign", result.Name)
    mockRepo.AssertExpectations(t)
}
```

### Integration Testing

```go
// tests/integration/api/campaigns_test.go
package api_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "ads-backend/internal/routes"
)

func TestCampaignsAPI_CreateCampaign(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    // Setup Gin router
    gin.SetMode(gin.TestMode)
    router := gin.New()
    
    // Setup routes
    campaignController := controllers.NewCampaignController(/* dependencies */)
    routes.SetupCampaignRoutes(router.Group("/api"), campaignController)
    
    // Prepare request
    campaign := map[string]interface{}{
        "name":           "Test Campaign",
        "budget_amount":  1000.00,
        "daily_budget":   50.00,
        "objective":      "conversions",
    }
    
    body, _ := json.Marshal(campaign)
    req := httptest.NewRequest("POST", "/api/campaigns", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+getTestJWTToken())
    
    // Execute request
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assert response
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, "Test Campaign", response["data"].(map[string]interface{})["name"])
}
```

## Deployment & DevOps

### Dockerfile for Go Application

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
```

### Docker Compose for Development

```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GO_ENV=development
      - DATABASE_URL=postgres://postgres:password@postgres:5432/ads_platform?sslmode=disable
      - REDIS_URL=redis://redis:6379
    volumes:
      - .:/app
      - /app/vendor
    depends_on:
      - postgres
      - redis
    command: air

  postgres:
    image: postgres:14-alpine
    environment:
      POSTGRES_DB: ads_platform
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  migrate:
    image: migrate/migrate
    networks:
      - default
    volumes:
      - ./migrations:/migrations
    command: ["-path", "/migrations", "-database", "postgres://postgres:password@postgres:5432/ads_platform?sslmode=disable", "up"]
    depends_on:
      - postgres

volumes:
  postgres_data:
```

### GitHub Actions CI/CD

```yaml
name: Go CI/CD

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:14
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: test_db
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21

    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Install dependencies
      run: go mod download

    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/test_db?sslmode=disable

    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out

  build:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21

    - name: Build
      run: go build -v ./cmd/server

    - name: Build Docker image
      run: docker build -t ads-backend .

    - name: Deploy to production
      run: |
        # Deployment script here
        echo "Deploying to production..."
```

## Environment Configuration

### Environment Variables

```bash
# Server Configuration
GO_ENV=production
PORT=8080
API_BASE_URL=https://api.ads-platform.com

# Database Configuration
DATABASE_URL=postgres://username:password@localhost:5432/ads_platform?sslmode=disable
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=25
DB_MAX_LIFETIME=5m

# Redis Configuration
REDIS_URL=redis://localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRES_IN=15m
REFRESH_TOKEN_SECRET=your-refresh-token-secret
REFRESH_TOKEN_EXPIRES_IN=168h

# Google APIs
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_ADS_DEVELOPER_TOKEN=your-developer-token
GEMINI_API_KEY=your-gemini-api-key
IMAGEN_API_KEY=your-imagen-api-key

# Email Service
EMAIL_SERVICE_API_KEY=your-email-service-key
EMAIL_FROM_ADDRESS=noreply@ads-platform.com

# Logging
LOG_LEVEL=info
LOG_FORMAT=json

# Rate Limiting
RATE_LIMIT_REQUESTS_PER_MINUTE=60
RATE_LIMIT_BURST=10
```

### Configuration Management

```go
// internal/config/config.go
package config

import (
    "log"
    "os"
    "strconv"
    "time"
    "github.com/joho/godotenv"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
    JWT      JWTConfig
    Google   GoogleConfig
    Email    EmailConfig
    Logging  LoggingConfig
}

type ServerConfig struct {
    Port       string
    Env        string
    BaseURL    string
}

type DatabaseConfig struct {
    URL          string
    MaxOpenConns int
    MaxIdleConns int
    MaxLifetime  time.Duration
}

type JWTConfig struct {
    Secret               string
    ExpiresIn            time.Duration
    RefreshSecret        string
    RefreshExpiresIn     time.Duration
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        Server: ServerConfig{
            Port:    getEnv("PORT", "8080"),
            Env:     getEnv("GO_ENV", "development"),
            BaseURL: getEnv("API_BASE_URL", "http://localhost:8080"),
        },
        Database: DatabaseConfig{
            URL:          getEnv("DATABASE_URL", ""),
            MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
            MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
            MaxLifetime:  getEnvAsDuration("DB_MAX_LIFETIME", "5m"),
        },
        JWT: JWTConfig{
            Secret:               getEnv("JWT_SECRET", ""),
            ExpiresIn:            getEnvAsDuration("JWT_EXPIRES_IN", "15m"),
            RefreshSecret:        getEnv("REFRESH_TOKEN_SECRET", ""),
            RefreshExpiresIn:     getEnvAsDuration("REFRESH_TOKEN_EXPIRES_IN", "168h"),
        },
        // ... other configs
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultVal
}

func getEnvAsDuration(name string, defaultVal string) time.Duration {
    valueStr := getEnv(name, defaultVal)
    if value, err := time.ParseDuration(valueStr); err == nil {
        return value
    }
    duration, _ := time.ParseDuration(defaultVal)
    return duration
}
```

## Frontend Integration

### API Client Updates for Go Backend

The frontend will need to update API calls to work with the new Go backend:

```typescript
// lib/api-client.ts (Frontend)
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const apiClient = axios.create({
  baseURL: `${API_BASE_URL}/api`,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add JWT token
apiClient.interceptors.request.use((config) => {
  const token = getAccessToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor for token refresh
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Try to refresh token
      const refreshed = await refreshAccessToken();
      if (refreshed) {
        // Retry original request
        return apiClient.request(error.config);
      }
      // Redirect to login
      window.location.href = '/auth/login';
    }
    return Promise.reject(error);
  }
);
```

## Migration Checklist

### Pre-Migration
- [ ] Set up new Go repository with proper project structure
- [ ] Configure development environment (Go 1.21+, PostgreSQL, Redis)
- [ ] Set up database migrations with golang-migrate
- [ ] Create basic Gin server with middleware

### Core Migration
- [ ] Convert Next.js API routes to Gin controllers
- [ ] Migrate AI services (Gemini, Google Ads, Imagen) to Go
- [ ] Implement GORM models and repositories
- [ ] Set up JWT authentication with golang-jwt
- [ ] Configure CORS and security middleware

### Advanced Features
- [ ] Implement WebSocket server with gorilla/websocket
- [ ] Add Redis caching for AI API responses
- [ ] Set up background jobs for analytics sync
- [ ] Implement rate limiting and request validation

### Testing & Documentation
- [ ] Write comprehensive unit tests with testify
- [ ] Create integration tests with testcontainers
- [ ] Generate Swagger documentation
- [ ] Set up test database and fixtures

### Frontend Updates
- [ ] Update API client to work with Go backend
- [ ] Modify authentication flow for JWT tokens
- [ ] Update WebSocket connection for Go server
- [ ] Test end-to-end functionality

### Deployment
- [ ] Create Docker configuration for Go application
- [ ] Set up CI/CD pipeline with GitHub Actions
- [ ] Configure production environment variables
- [ ] Deploy to staging and production environments

### Performance & Monitoring
- [ ] Set up logging with structured logging
- [ ] Implement health checks and metrics
- [ ] Configure monitoring and alerting
- [ ] Perform load testing and optimization

---

This comprehensive Go backend documentation provides all the necessary details for successfully migrating from the Next.js full-stack application to a dedicated Go API server with high performance, scalability, and maintainability.

adust09/vibe-marketing-backend
adust09/vibe-marketing-backend
