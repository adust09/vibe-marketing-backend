package routes

import (
	"net/http"

	"ads-backend/internal/controllers"
	"ads-backend/internal/services"
	"ads-backend/internal/repositories"
	"ads-backend/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// Setup route groups
	// SetupAuthRoutes(router)
	// SetupCampaignRoutes(router)
	// SetupAnalyticsRoutes(router)
	// SetupAIRoutes(router)
	// SetupUserRoutes(router)
	SetupDemographicsRoutes(router, db, cfg)
}

func SetupDemographicsRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	// Initialize dependencies
	demographicsRepo := repositories.NewDemographicsRepository(db)
	demographicsService := services.NewDemographicsService(demographicsRepo, cfg)
	demographicsController := controllers.NewDemographicsController(demographicsService)

	// Demographics routes
	demographics := router.Group("/demographics")
	{
		// Get user demographics
		demographics.GET("/users/:user_id", demographicsController.GetUserDemographics)
		
		// Update user demographics from Google Ads API
		demographics.PUT("/users/:user_id", demographicsController.UpdateUserDemographics)
		
		// Get demographics summary (age and gender distribution)
		demographics.GET("/summary", demographicsController.GetDemographicsSummary)
		
		// Get performance metrics by demographics
		demographics.GET("/performance", demographicsController.GetPerformanceByDemographics)
		
		// Admin routes for refreshing demographics data
		demographics.POST("/refresh/all", demographicsController.RefreshAllDemographics)
		demographics.POST("/refresh/stale", demographicsController.RefreshStaleRecords)
	}
}