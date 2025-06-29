package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/models"
	"ads-backend/internal/services"
	"ads-backend/internal/utils"
)

type CampaignController struct {
	campaignService *services.CampaignService
}

func NewCampaignController(campaignService *services.CampaignService) *CampaignController {
	return &CampaignController{
		campaignService: campaignService,
	}
}

func (cc *CampaignController) GetCampaigns(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	campaigns, err := cc.campaignService.GetCampaignsByUserID(userID.(uuid.UUID))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch campaigns")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Campaigns retrieved successfully", campaigns)
}

func (cc *CampaignController) GetCampaign(c *gin.Context) {
	campaignID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid campaign ID")
		return
	}

	campaign, err := cc.campaignService.GetCampaignByID(campaignID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Campaign not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Campaign retrieved successfully", campaign)
}

func (cc *CampaignController) UpdateCampaignCPC(c *gin.Context) {
	campaignID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid campaign ID")
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

	campaign, err := cc.campaignService.UpdateCampaignCPC(campaignID, cpcRequest.CPC, cpcRequest.AverageCPC, cpcRequest.MaxCPC)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update campaign CPC")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Campaign CPC updated successfully", campaign)
}

func (cc *CampaignController) RefreshCampaignCPC(c *gin.Context) {
	campaignID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid campaign ID")
		return
	}

	campaign, err := cc.campaignService.RefreshCampaignCPCFromGoogleAds(campaignID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh campaign CPC from Google Ads")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Campaign CPC refreshed successfully", campaign)
}