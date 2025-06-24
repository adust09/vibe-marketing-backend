package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/models"
	"ads-backend/internal/repositories"
)

type KeywordService struct {
	keywordRepo      *repositories.KeywordRepository
	googleAdsService *GoogleAdsService
}

func NewKeywordService(keywordRepo *repositories.KeywordRepository, googleAdsService *GoogleAdsService) *KeywordService {
	return &KeywordService{
		keywordRepo:      keywordRepo,
		googleAdsService: googleAdsService,
	}
}

func (ks *KeywordService) GetKeywordsByAdGroupID(adGroupID uuid.UUID) ([]models.Keyword, error) {
	keywords, err := ks.keywordRepo.GetByAdGroupID(adGroupID)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

func (ks *KeywordService) GetKeywordByID(keywordID uuid.UUID) (*models.Keyword, error) {
	keyword, err := ks.keywordRepo.GetByID(keywordID)
	if err != nil {
		return nil, err
	}
	return keyword, nil
}

func (ks *KeywordService) UpdateKeywordCPC(keywordID uuid.UUID, cpc, averageCPC, maxCPC *decimal.Decimal) (*models.Keyword, error) {
	keyword, err := ks.keywordRepo.GetByID(keywordID)
	if err != nil {
		return nil, err
	}

	keyword.CPC = cpc
	keyword.AverageCPC = averageCPC
	keyword.MaxCPC = maxCPC

	updatedKeyword, err := ks.keywordRepo.Update(keyword)
	if err != nil {
		return nil, err
	}

	return updatedKeyword, nil
}

func (ks *KeywordService) RefreshKeywordCPCFromGoogleAds(keywordID uuid.UUID) (*models.Keyword, error) {
	keyword, err := ks.keywordRepo.GetByID(keywordID)
	if err != nil {
		return nil, err
	}

	if keyword.GoogleAdsKeywordID == nil {
		return nil, errors.New("keyword is not linked to Google Ads")
	}

	cpcData, err := ks.googleAdsService.GetKeywordCPC(*keyword.GoogleAdsKeywordID)
	if err != nil {
		return nil, err
	}

	keyword.CPC = cpcData.CPC
	keyword.AverageCPC = cpcData.AverageCPC
	keyword.MaxCPC = cpcData.MaxCPC
	keyword.QualityScore = cpcData.QualityScore
	keyword.Impressions = cpcData.Impressions
	keyword.Clicks = cpcData.Clicks
	keyword.Cost = cpcData.Cost

	updatedKeyword, err := ks.keywordRepo.Update(keyword)
	if err != nil {
		return nil, err
	}

	return updatedKeyword, nil
}