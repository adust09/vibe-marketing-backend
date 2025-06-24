package models

import (
	"time"

	"github.com/google/uuid"
)

type AgeRange string

const (
	AgeRange18To24 AgeRange = "18-24"
	AgeRange25To34 AgeRange = "25-34"
	AgeRange35To44 AgeRange = "35-44"
	AgeRange45To54 AgeRange = "45-54"
	AgeRange55To64 AgeRange = "55-64"
	AgeRange65Plus AgeRange = "65+"
	AgeRangeUnknown AgeRange = "unknown"
)

type Gender string

const (
	GenderMale     Gender = "male"
	GenderFemale   Gender = "female"
	GenderOther    Gender = "other"
	GenderUnknown  Gender = "unknown"
)

type UserDemographics struct {
	BaseModel
	UserID              uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User                User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AgeRange            AgeRange  `gorm:"type:varchar(20)" json:"age_range"`
	Gender              Gender    `gorm:"type:varchar(20)" json:"gender"`
	LastUpdatedFromAPI  time.Time `json:"last_updated_from_api"`
	DataSource          string    `gorm:"type:varchar(50);default:'google_ads'" json:"data_source"`
	Confidence          *float64  `json:"confidence,omitempty"`
	PrivacyCompliant    bool      `gorm:"default:true" json:"privacy_compliant"`
}

type DemographicsSummary struct {
	AgeDistribution    map[AgeRange]int `json:"age_distribution"`
	GenderDistribution map[Gender]int   `json:"gender_distribution"`
	TotalUsers         int              `json:"total_users"`
	LastUpdated        time.Time        `json:"last_updated"`
}

type DemographicsPerformance struct {
	AgeRange         AgeRange `json:"age_range"`
	Gender           Gender   `json:"gender"`
	UserCount        int      `json:"user_count"`
	ClickThroughRate float64  `json:"click_through_rate"`
	ConversionRate   float64  `json:"conversion_rate"`
	AverageCPC       float64  `json:"average_cpc"`
	ROAS             float64  `json:"roas"`
}