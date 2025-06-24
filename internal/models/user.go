package models

type User struct {
	BaseModel
	Email               string     `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash        string     `gorm:"not null" json:"-"`
	FirstName           *string    `json:"first_name"`
	LastName            *string    `json:"last_name"`
	Role                string     `gorm:"default:user" json:"role"`
	IsActive            bool       `gorm:"default:true" json:"is_active"`
	EmailVerified       bool       `gorm:"default:false" json:"email_verified"`
	GoogleAdsCustomerID *string    `json:"google_ads_customer_id"`
	Campaigns           []Campaign `gorm:"foreignKey:UserID" json:"campaigns,omitempty"`
}