package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type Campaign struct {
	BaseModel
	UserID              uuid.UUID           `gorm:"type:uuid;not null" json:"user_id"`
	Name                string              `gorm:"not null" json:"name"`
	Status              string              `gorm:"default:draft" json:"status"`
	BudgetAmount        decimal.Decimal     `gorm:"type:decimal(12,2);not null" json:"budget_amount"`
	DailyBudget         decimal.Decimal     `gorm:"type:decimal(12,2);not null" json:"daily_budget"`
	StartDate           *time.Time          `gorm:"type:date" json:"start_date"`
	EndDate             *time.Time          `gorm:"type:date" json:"end_date"`
	Objective           string              `gorm:"not null" json:"objective"`
	TargetAudience      datatypes.JSON      `json:"target_audience"`
	GoogleAdsCampaignID *string             `json:"google_ads_campaign_id"`
	Platform            string              `gorm:"default:google" json:"platform"`
	User                User                `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AdGroups            []AdGroup           `gorm:"foreignKey:CampaignID" json:"ad_groups,omitempty"`
	Analytics           []CampaignAnalytics `gorm:"foreignKey:CampaignID" json:"analytics,omitempty"`
}