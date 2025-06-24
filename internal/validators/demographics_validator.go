package validators

import (
	"strconv"

	"ads-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DemographicsValidator struct{}

func NewDemographicsValidator() *DemographicsValidator {
	return &DemographicsValidator{}
}

func (v *DemographicsValidator) ValidateUserID(c *gin.Context) (uuid.UUID, error) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func (v *DemographicsValidator) ValidateRefreshParams(c *gin.Context) (int, error) {
	olderThanHours := 24
	if hoursParam := c.DefaultQuery("older_than_hours", "24"); hoursParam != "" {
		hours, err := strconv.Atoi(hoursParam)
		if err != nil || hours < 1 || hours > 8760 { // max 1 year
			return 0, err
		}
		olderThanHours = hours
	}
	return olderThanHours, nil
}

func (v *DemographicsValidator) ValidateAgeRange(ageRange string) bool {
	validAgeRanges := []models.AgeRange{
		models.AgeRange18To24,
		models.AgeRange25To34,
		models.AgeRange35To44,
		models.AgeRange45To54,
		models.AgeRange55To64,
		models.AgeRange65Plus,
		models.AgeRangeUnknown,
	}
	
	for _, validRange := range validAgeRanges {
		if string(validRange) == ageRange {
			return true
		}
	}
	return false
}

func (v *DemographicsValidator) ValidateGender(gender string) bool {
	validGenders := []models.Gender{
		models.GenderMale,
		models.GenderFemale,
		models.GenderOther,
		models.GenderUnknown,
	}
	
	for _, validGender := range validGenders {
		if string(validGender) == gender {
			return true
		}
	}
	return false
}

func (v *DemographicsValidator) ValidateConfidence(confidence *float64) bool {
	if confidence == nil {
		return true
	}
	return *confidence >= 0.0 && *confidence <= 1.0
}