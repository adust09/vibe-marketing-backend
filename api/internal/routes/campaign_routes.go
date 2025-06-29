package routes

import (
	"github.com/gin-gonic/gin"

	"ads-backend/internal/controllers"
)

func SetupCampaignRoutes(router *gin.RouterGroup, campaignController *controllers.CampaignController) {
	campaigns := router.Group("/campaigns")
	{
		campaigns.GET("/", campaignController.GetCampaigns)
		campaigns.GET("/:id", campaignController.GetCampaign)
		campaigns.PUT("/:id/cpc", campaignController.UpdateCampaignCPC)
		campaigns.POST("/:id/refresh-cpc", campaignController.RefreshCampaignCPC)
	}
}