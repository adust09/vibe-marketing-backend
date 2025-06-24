package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Keyword struct {
	BaseModel
	AdGroupID            uuid.UUID        `gorm:"type:uuid;not null" json:"ad_group_id"`
	Text                 string           `gorm:"not null" json:"text"`
	MatchType            string           `gorm:"not null" json:"match_type"`
	Status               string           `gorm:"default:active" json:"status"`
	GoogleAdsKeywordID   *string          `json:"google_ads_keyword_id"`
	CPC                  *decimal.Decimal `gorm:"type:decimal(10,4)" json:"cpc"`
	AverageCPC           *decimal.Decimal `gorm:"type:decimal(10,4)" json:"average_cpc"`
	MaxCPC               *decimal.Decimal `gorm:"type:decimal(10,4)" json:"max_cpc"`
	QualityScore         *int             `json:"quality_score"`
	Impressions          *int64           `json:"impressions"`
	Clicks               *int64           `json:"clicks"`
	Cost                 *decimal.Decimal `gorm:"type:decimal(12,2)" json:"cost"`
	AdGroup              AdGroup          `gorm:"foreignKey:AdGroupID" json:"ad_group,omitempty"`
}