package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"ads-backend/internal/models"
)

type AdGroupRepository struct {
	db *gorm.DB
}

func NewAdGroupRepository(db *gorm.DB) *AdGroupRepository {
	return &AdGroupRepository{db: db}
}

func (agr *AdGroupRepository) GetByCampaignID(campaignID uuid.UUID) ([]models.AdGroup, error) {
	var adGroups []models.AdGroup
	err := agr.db.Where("campaign_id = ?", campaignID).Find(&adGroups).Error
	return adGroups, err
}

func (agr *AdGroupRepository) GetByID(adGroupID uuid.UUID) (*models.AdGroup, error) {
	var adGroup models.AdGroup
	err := agr.db.Where("id = ?", adGroupID).First(&adGroup).Error
	if err != nil {
		return nil, err
	}
	return &adGroup, nil
}

func (agr *AdGroupRepository) Update(adGroup *models.AdGroup) (*models.AdGroup, error) {
	err := agr.db.Save(adGroup).Error
	if err != nil {
		return nil, err
	}
	return adGroup, nil
}

func (agr *AdGroupRepository) Create(adGroup *models.AdGroup) (*models.AdGroup, error) {
	err := agr.db.Create(adGroup).Error
	if err != nil {
		return nil, err
	}
	return adGroup, nil
}

func (agr *AdGroupRepository) Delete(adGroupID uuid.UUID) error {
	return agr.db.Delete(&models.AdGroup{}, adGroupID).Error
}