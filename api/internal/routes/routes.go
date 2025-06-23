package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup) {
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
}