package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ads-backend/internal/config"
	"ads-backend/internal/controllers"
	"ads-backend/internal/repositories"
	"ads-backend/internal/services"
)

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	// Base API info
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ads Backend API v1",
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// Initialize repositories
	campaignRepo := repositories.NewCampaignRepository(db)
	adGroupRepo := repositories.NewAdGroupRepository(db)
	keywordRepo := repositories.NewKeywordRepository(db)

	// Initialize services
	googleAdsService := services.NewGoogleAdsService(cfg)
	campaignService := services.NewCampaignService(campaignRepo, googleAdsService)
	adGroupService := services.NewAdGroupService(adGroupRepo, googleAdsService)
	keywordService := services.NewKeywordService(keywordRepo, googleAdsService)

	// Initialize controllers
	campaignController := controllers.NewCampaignController(campaignService)
	adGroupController := controllers.NewAdGroupController(adGroupService)
	keywordController := controllers.NewKeywordController(keywordService)

	// Setup route groups
	// SetupAuthRoutes(router)
	SetupCampaignRoutes(router, campaignController)
	SetupAdGroupRoutes(router, adGroupController)
	SetupKeywordRoutes(router, keywordController)
	// SetupAnalyticsRoutes(router)
	// SetupAIRoutes(router)
	// SetupUserRoutes(router)
}