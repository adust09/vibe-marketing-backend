package controllers

import (
	"net/http"
	"time"

	"ads-backend/internal/services"
	"ads-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DemographicsController struct {
	service services.DemographicsService
}

func NewDemographicsController(service services.DemographicsService) *DemographicsController {
	return &DemographicsController{
		service: service,
	}
}

func (dc *DemographicsController) GetUserDemographics(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	demographics, err := dc.service.GetUserDemographics(userID)
	if err != nil {
		if err.Error() == "record not found" {
			utils.ErrorResponse(c, http.StatusNotFound, "User demographics not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get user demographics")
		return
	}

	utils.SuccessResponse(c, "User demographics retrieved successfully", demographics)
}

func (dc *DemographicsController) UpdateUserDemographics(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	err = dc.service.UpdateUserDemographics(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user demographics")
		return
	}

	utils.SuccessResponse(c, "User demographics updated successfully", nil)
}

func (dc *DemographicsController) GetDemographicsSummary(c *gin.Context) {
	summary, err := dc.service.GetDemographicsSummary()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get demographics summary")
		return
	}

	utils.SuccessResponse(c, "Demographics summary retrieved successfully", summary)
}

func (dc *DemographicsController) GetPerformanceByDemographics(c *gin.Context) {
	performance, err := dc.service.GetPerformanceByDemographics()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get performance by demographics")
		return
	}

	utils.SuccessResponse(c, "Demographics performance retrieved successfully", performance)
}

func (dc *DemographicsController) RefreshAllDemographics(c *gin.Context) {
	ctx := c.Request.Context()
	
	err := dc.service.RefreshAllDemographics(ctx)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh demographics")
		return
	}

	utils.SuccessResponse(c, "Demographics refresh started successfully", nil)
}

func (dc *DemographicsController) RefreshStaleRecords(c *gin.Context) {
	ctx := c.Request.Context()
	
	olderThanHours := 24
	if hoursParam := c.DefaultQuery("older_than_hours", "24"); hoursParam != "" {
		if hours, err := time.ParseDuration(hoursParam + "h"); err == nil {
			olderThanHours = int(hours.Hours())
		}
	}
	
	olderThan := time.Duration(olderThanHours) * time.Hour
	
	err := dc.service.RefreshStaleRecords(ctx, olderThan)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to refresh stale demographics records")
		return
	}

	utils.SuccessResponse(c, "Stale demographics records refreshed successfully", nil)
}