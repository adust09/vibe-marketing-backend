package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"ads-backend/internal/models"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{db: db}
}

func (cr *CampaignRepository) GetByUserID(userID uuid.UUID) ([]models.Campaign, error) {
	var campaigns []models.Campaign
	err := cr.db.Where("user_id = ?", userID).Find(&campaigns).Error
	return campaigns, err
}

func (cr *CampaignRepository) GetByID(campaignID uuid.UUID) (*models.Campaign, error) {
	var campaign models.Campaign
	err := cr.db.Where("id = ?", campaignID).First(&campaign).Error
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (cr *CampaignRepository) Update(campaign *models.Campaign) (*models.Campaign, error) {
	err := cr.db.Save(campaign).Error
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (cr *CampaignRepository) Create(campaign *models.Campaign) (*models.Campaign, error) {
	err := cr.db.Create(campaign).Error
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (cr *CampaignRepository) Delete(campaignID uuid.UUID) error {
	return cr.db.Delete(&models.Campaign{}, campaignID).Error
}