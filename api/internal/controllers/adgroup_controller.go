package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/services"
	"ads-backend/internal/utils"
)

type AdGroupController struct {
	adGroupService *services.AdGroupService
}

func NewAdGroupController(adGroupService *services.AdGroupService) *AdGroupController {
	return &AdGroupController{
		adGroupService: adGroupService,
	}
}

func (agc *AdGroupController) GetAdGroups(c *gin.Context) {
	campaignID, err := uuid.Parse(c.Param("campaign_id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid campaign ID")
		return
	}

	adGroups, err := agc.adGroupService.GetAdGroupsByCampaignID(campaignID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch ad groups")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Ad groups retrieved successfully", adGroups)
}

func (agc *AdGroupController) GetAdGroup(c *gin.Context) {
	adGroupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ad group ID")
		return
	}

	adGroup, err := agc.adGroupService.GetAdGroupByID(adGroupID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Ad group not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Ad group retrieved successfully", adGroup)
}

func (agc *AdGroupController) UpdateAdGroupCPC(c *gin.Context) {
	adGroupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ad group ID")
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

	adGroup, err := agc.adGroupService.UpdateAdGroupCPC(adGroupID, cpcRequest.CPC, cpcRequest.AverageCPC, cpcRequest.MaxCPC)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update ad group CPC")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Ad group CPC updated successfully", adGroup)
}

func (agc *AdGroupController) RefreshAdGroupCPC(c *gin.Context) {
	adGroupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ad group ID")
		return
	}

	adGroup, err := agc.adGroupService.RefreshAdGroupCPCFromGoogleAds(adGroupID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh ad group CPC from Google Ads")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Ad group CPC refreshed successfully", adGroup)
}