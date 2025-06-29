package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/models"
	"ads-backend/internal/repositories"
)

type AdGroupService struct {
	adGroupRepo      *repositories.AdGroupRepository
	googleAdsService *GoogleAdsService
}

func NewAdGroupService(adGroupRepo *repositories.AdGroupRepository, googleAdsService *GoogleAdsService) *AdGroupService {
	return &AdGroupService{
		adGroupRepo:      adGroupRepo,
		googleAdsService: googleAdsService,
	}
}

func (ags *AdGroupService) GetAdGroupsByCampaignID(campaignID uuid.UUID) ([]models.AdGroup, error) {
	adGroups, err := ags.adGroupRepo.GetByCampaignID(campaignID)
	if err != nil {
		return nil, err
	}
	return adGroups, nil
}

func (ags *AdGroupService) GetAdGroupByID(adGroupID uuid.UUID) (*models.AdGroup, error) {
	adGroup, err := ags.adGroupRepo.GetByID(adGroupID)
	if err != nil {
		return nil, err
	}
	return adGroup, nil
}

func (ags *AdGroupService) UpdateAdGroupCPC(adGroupID uuid.UUID, cpc, averageCPC, maxCPC *decimal.Decimal) (*models.AdGroup, error) {
	adGroup, err := ags.adGroupRepo.GetByID(adGroupID)
	if err != nil {
		return nil, err
	}

	adGroup.CPC = cpc
	adGroup.AverageCPC = averageCPC
	adGroup.MaxCPC = maxCPC

	updatedAdGroup, err := ags.adGroupRepo.Update(adGroup)
	if err != nil {
		return nil, err
	}

	return updatedAdGroup, nil
}

func (ags *AdGroupService) RefreshAdGroupCPCFromGoogleAds(adGroupID uuid.UUID) (*models.AdGroup, error) {
	adGroup, err := ags.adGroupRepo.GetByID(adGroupID)
	if err != nil {
		return nil, err
	}

	if adGroup.GoogleAdsAdGroupID == nil {
		return nil, errors.New("ad group is not linked to Google Ads")
	}

	cpcData, err := ags.googleAdsService.GetAdGroupCPC(*adGroup.GoogleAdsAdGroupID)
	if err != nil {
		return nil, err
	}

	adGroup.CPC = cpcData.CPC
	adGroup.AverageCPC = cpcData.AverageCPC
	adGroup.MaxCPC = cpcData.MaxCPC

	updatedAdGroup, err := ags.adGroupRepo.Update(adGroup)
	if err != nil {
		return nil, err
	}

	return updatedAdGroup, nil
}