package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type AdGroup struct {
	BaseModel
	CampaignID         uuid.UUID        `gorm:"type:uuid;not null" json:"campaign_id"`
	Name               string           `gorm:"not null" json:"name"`
	Status             string           `gorm:"default:active" json:"status"`
	Targeting          datatypes.JSON   `json:"targeting"`
	GoogleAdsAdGroupID *string          `json:"google_ads_ad_group_id"`
	CPC                *decimal.Decimal `gorm:"type:decimal(10,4)" json:"cpc"`
	AverageCPC         *decimal.Decimal `gorm:"type:decimal(10,4)" json:"average_cpc"`
	MaxCPC             *decimal.Decimal `gorm:"type:decimal(10,4)" json:"max_cpc"`
	Campaign           Campaign         `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
	Keywords           []Keyword        `gorm:"foreignKey:AdGroupID" json:"keywords,omitempty"`
}