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
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port    string
	Env     string
	BaseURL string
}

type DatabaseConfig struct {
	URL          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

type RedisConfig struct {
	URL      string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret           string
	ExpiresIn        time.Duration
	RefreshSecret    string
	RefreshExpiresIn time.Duration
}

type GoogleConfig struct {
	ClientID          string
	ClientSecret      string
	AdsDeveloperToken string
	GeminiAPIKey      string
	ImagenAPIKey      string
}

type EmailConfig struct {
	ServiceAPIKey string
	FromAddress   string
}

type LoggingConfig struct {
	Level  string
	Format string
}

type RateLimitConfig struct {
	RequestsPerMinute int
	Burst            int
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
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
		Redis: RedisConfig{
			URL:      getEnv("REDIS_URL", "redis://localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret:           getEnv("JWT_SECRET", ""),
			ExpiresIn:        getEnvAsDuration("JWT_EXPIRES_IN", "15m"),
			RefreshSecret:    getEnv("REFRESH_TOKEN_SECRET", ""),
			RefreshExpiresIn: getEnvAsDuration("REFRESH_TOKEN_EXPIRES_IN", "168h"),
		},
		Google: GoogleConfig{
			ClientID:          getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret:      getEnv("GOOGLE_CLIENT_SECRET", ""),
			AdsDeveloperToken: getEnv("GOOGLE_ADS_DEVELOPER_TOKEN", ""),
			GeminiAPIKey:      getEnv("GEMINI_API_KEY", ""),
			ImagenAPIKey:      getEnv("IMAGEN_API_KEY", ""),
		},
		Email: EmailConfig{
			ServiceAPIKey: getEnv("EMAIL_SERVICE_API_KEY", ""),
			FromAddress:   getEnv("EMAIL_FROM_ADDRESS", "noreply@ads-platform.com"),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		RateLimit: RateLimitConfig{
			RequestsPerMinute: getEnvAsInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 60),
			Burst:            getEnvAsInt("RATE_LIMIT_BURST", 10),
		},
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