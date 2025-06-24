package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"ads-backend/internal/models"
	"ads-backend/internal/repositories"
)

type CampaignService struct {
	campaignRepo   *repositories.CampaignRepository
	googleAdsService *GoogleAdsService
}

func NewCampaignService(campaignRepo *repositories.CampaignRepository, googleAdsService *GoogleAdsService) *CampaignService {
	return &CampaignService{
		campaignRepo:   campaignRepo,
		googleAdsService: googleAdsService,
	}
}

func (cs *CampaignService) GetCampaignsByUserID(userID uuid.UUID) ([]models.Campaign, error) {
	campaigns, err := cs.campaignRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (cs *CampaignService) GetCampaignByID(campaignID uuid.UUID) (*models.Campaign, error) {
	campaign, err := cs.campaignRepo.GetByID(campaignID)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (cs *CampaignService) UpdateCampaignCPC(campaignID uuid.UUID, cpc, averageCPC, maxCPC *decimal.Decimal) (*models.Campaign, error) {
	campaign, err := cs.campaignRepo.GetByID(campaignID)
	if err != nil {
		return nil, err
	}

	campaign.CPC = cpc
	campaign.AverageCPC = averageCPC
	campaign.MaxCPC = maxCPC

	updatedCampaign, err := cs.campaignRepo.Update(campaign)
	if err != nil {
		return nil, err
	}

	return updatedCampaign, nil
}

func (cs *CampaignService) RefreshCampaignCPCFromGoogleAds(campaignID uuid.UUID) (*models.Campaign, error) {
	campaign, err := cs.campaignRepo.GetByID(campaignID)
	if err != nil {
		return nil, err
	}

	if campaign.GoogleAdsCampaignID == nil {
		return nil, errors.New("campaign is not linked to Google Ads")
	}

	cpcData, err := cs.googleAdsService.GetCampaignCPC(*campaign.GoogleAdsCampaignID)
	if err != nil {
		return nil, err
	}

	campaign.CPC = cpcData.CPC
	campaign.AverageCPC = cpcData.AverageCPC
	campaign.MaxCPC = cpcData.MaxCPC

	updatedCampaign, err := cs.campaignRepo.Update(campaign)
	if err != nil {
		return nil, err
	}

	return updatedCampaign, nil
}