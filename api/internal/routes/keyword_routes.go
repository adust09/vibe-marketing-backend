package routes

import (
	"github.com/gin-gonic/gin"

	"ads-backend/internal/controllers"
)

func SetupKeywordRoutes(router *gin.RouterGroup, keywordController *controllers.KeywordController) {
	adGroups := router.Group("/adgroups")
	{
		adGroups.GET("/:adgroup_id/keywords", keywordController.GetKeywords)
	}

	keywords := router.Group("/keywords")
	{
		keywords.GET("/:id", keywordController.GetKeyword)
		keywords.PUT("/:id/cpc", keywordController.UpdateKeywordCPC)
		keywords.POST("/:id/refresh-cpc", keywordController.RefreshKeywordCPC)
	}
}