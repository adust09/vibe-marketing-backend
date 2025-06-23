package config

import (
	"log"

	"ads-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase(cfg *Config) *gorm.DB {
	var gormLogger logger.Interface

	if cfg.Server.Env == "development" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.MaxLifetime)

	// Auto-migrate models (for development)
	if cfg.Server.Env == "development" {
		err = db.AutoMigrate(
			&models.User{},
			&models.Campaign{},
			&models.AdGroup{},
			&models.CampaignAnalytics{},
		)
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	}

	log.Println("Database connected successfully")
	return db
}