package models

import (
	"time"

	"github.com/google/uuid"
)

type CampaignAnalytics struct {
	BaseModel
	CampaignID    uuid.UUID `gorm:"type:uuid;not null" json:"campaign_id"`
	Date          time.Time `gorm:"type:date;not null" json:"date"`
	Impressions   int       `gorm:"default:0" json:"impressions"`
	Clicks        int       `gorm:"default:0" json:"clicks"`
	Conversions   int       `gorm:"default:0" json:"conversions"`
	CostMicros    int64     `gorm:"default:0" json:"cost_micros"`
	RevenueMicros *int64    `json:"revenue_micros"`
	Campaign      Campaign  `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
}