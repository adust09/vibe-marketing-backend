package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"ads-backend/internal/models"
)

type KeywordRepository struct {
	db *gorm.DB
}

func NewKeywordRepository(db *gorm.DB) *KeywordRepository {
	return &KeywordRepository{db: db}
}

func (kr *KeywordRepository) GetByAdGroupID(adGroupID uuid.UUID) ([]models.Keyword, error) {
	var keywords []models.Keyword
	err := kr.db.Where("ad_group_id = ?", adGroupID).Find(&keywords).Error
	return keywords, err
}

func (kr *KeywordRepository) GetByID(keywordID uuid.UUID) (*models.Keyword, error) {
	var keyword models.Keyword
	err := kr.db.Where("id = ?", keywordID).First(&keyword).Error
	if err != nil {
		return nil, err
	}
	return &keyword, nil
}

func (kr *KeywordRepository) Update(keyword *models.Keyword) (*models.Keyword, error) {
	err := kr.db.Save(keyword).Error
	if err != nil {
		return nil, err
	}
	return keyword, nil
}

func (kr *KeywordRepository) Create(keyword *models.Keyword) (*models.Keyword, error) {
	err := kr.db.Create(keyword).Error
	if err != nil {
		return nil, err
	}
	return keyword, nil
}

func (kr *KeywordRepository) Delete(keywordID uuid.UUID) error {
	return kr.db.Delete(&models.Keyword{}, keywordID).Error
}