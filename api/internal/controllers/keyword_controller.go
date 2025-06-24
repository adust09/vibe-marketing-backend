package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/services"
	"ads-backend/internal/utils"
)

type KeywordController struct {
	keywordService *services.KeywordService
}

func NewKeywordController(keywordService *services.KeywordService) *KeywordController {
	return &KeywordController{
		keywordService: keywordService,
	}
}

func (kc *KeywordController) GetKeywords(c *gin.Context) {
	adGroupID, err := uuid.Parse(c.Param("adgroup_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ad group ID")
		return
	}

	keywords, err := kc.keywordService.GetKeywordsByAdGroupID(adGroupID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch keywords")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Keywords retrieved successfully", keywords)
}

func (kc *KeywordController) GetKeyword(c *gin.Context) {
	keywordID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid keyword ID")
		return
	}

	keyword, err := kc.keywordService.GetKeywordByID(keywordID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Keyword not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Keyword retrieved successfully", keyword)
}

func (kc *KeywordController) UpdateKeywordCPC(c *gin.Context) {
	keywordID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid keyword ID")
		return
	}

	var cpcRequest struct {
		CPC        *decimal.Decimal `json:"cpc"`
		AverageCPC *decimal.Decimal `json:"average_cpc"`
		MaxCPC     *decimal.Decimal `json:"max_cpc"`
	}

	if err := c.ShouldBindJSON(&cpcRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	keyword, err := kc.keywordService.UpdateKeywordCPC(keywordID, cpcRequest.CPC, cpcRequest.AverageCPC, cpcRequest.MaxCPC)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update keyword CPC")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Keyword CPC updated successfully", keyword)
}

func (kc *KeywordController) RefreshKeywordCPC(c *gin.Context) {
	keywordID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid keyword ID")
		return
	}

	keyword, err := kc.keywordService.RefreshKeywordCPCFromGoogleAds(keywordID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh keyword CPC from Google Ads")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Keyword CPC refreshed successfully", keyword)
}