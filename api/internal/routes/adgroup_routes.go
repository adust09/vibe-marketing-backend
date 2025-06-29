package routes

import (
	"github.com/gin-gonic/gin"

	"ads-backend/internal/controllers"
)

func SetupAdGroupRoutes(router *gin.RouterGroup, adGroupController *controllers.AdGroupController) {
	campaigns := router.Group("/campaigns")
	{
		campaigns.GET("/:campaign_id/adgroups", adGroupController.GetAdGroups)
	}

	adGroups := router.Group("/adgroups")
	{
		adGroups.GET("/:id", adGroupController.GetAdGroup)
		adGroups.PUT("/:id/cpc", adGroupController.UpdateAdGroupCPC)
		adGroups.POST("/:id/refresh-cpc", adGroupController.RefreshAdGroupCPC)
	}
}