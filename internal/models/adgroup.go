package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AdGroup struct {
	BaseModel
	CampaignID         uuid.UUID      `gorm:"type:uuid;not null" json:"campaign_id"`
	Name               string         `gorm:"not null" json:"name"`
	Status             string         `gorm:"default:active" json:"status"`
	Targeting          datatypes.JSON `json:"targeting"`
	GoogleAdsAdGroupID *string        `json:"google_ads_ad_group_id"`
	Campaign           Campaign       `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
}