package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ads-backend/internal/models"
	"ads-backend/internal/services"
)

// Mock repository
type MockCampaignRepository struct {
	mock.Mock
}

func (m *MockCampaignRepository) GetByUserID(userID uuid.UUID) ([]models.Campaign, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) GetByID(campaignID uuid.UUID) (*models.Campaign, error) {
	args := m.Called(campaignID)
	return args.Get(0).(*models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) Update(campaign *models.Campaign) (*models.Campaign, error) {
	args := m.Called(campaign)
	return args.Get(0).(*models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) Create(campaign *models.Campaign) (*models.Campaign, error) {
	args := m.Called(campaign)
	return args.Get(0).(*models.Campaign), args.Error(1)
}

func (m *MockCampaignRepository) Delete(campaignID uuid.UUID) error {
	args := m.Called(campaignID)
	return args.Error(0)
}

// Mock Google Ads service
type MockGoogleAdsService struct {
	mock.Mock
}

func (m *MockGoogleAdsService) GetCampaignCPC(campaignID string) (*services.CPCData, error) {
	args := m.Called(campaignID)
	return args.Get(0).(*services.CPCData), args.Error(1)
}

func (m *MockGoogleAdsService) GetAdGroupCPC(adGroupID string) (*services.CPCData, error) {
	args := m.Called(adGroupID)
	return args.Get(0).(*services.CPCData), args.Error(1)
}

func (m *MockGoogleAdsService) GetKeywordCPC(keywordID string) (*services.CPCData, error) {
	args := m.Called(keywordID)
	return args.Get(0).(*services.CPCData), args.Error(1)
}

func TestCampaignService_UpdateCampaignCPC(t *testing.T) {
	// Setup
	mockRepo := new(MockCampaignRepository)
	mockGoogleAds := new(MockGoogleAdsService)
	campaignService := services.NewCampaignService(mockRepo, mockGoogleAds)

	campaignID := uuid.New()
	cpc := decimal.NewFromFloat(1.50)
	avgCPC := decimal.NewFromFloat(1.25)
	maxCPC := decimal.NewFromFloat(2.00)

	campaign := &models.Campaign{
		BaseModel: models.BaseModel{ID: campaignID},
		Name:      "Test Campaign",
	}

	updatedCampaign := &models.Campaign{
		BaseModel:  models.BaseModel{ID: campaignID},
		Name:       "Test Campaign",
		CPC:        &cpc,
		AverageCPC: &avgCPC,
		MaxCPC:     &maxCPC,
	}

	// Set expectations
	mockRepo.On("GetByID", campaignID).Return(campaign, nil)
	mockRepo.On("Update", mock.MatchedBy(func(c *models.Campaign) bool {
		return c.ID == campaignID && 
			   c.CPC.Equal(cpc) && 
			   c.AverageCPC.Equal(avgCPC) && 
			   c.MaxCPC.Equal(maxCPC)
	})).Return(updatedCampaign, nil)

	// Execute
	result, err := campaignService.UpdateCampaignCPC(campaignID, &cpc, &avgCPC, &maxCPC)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, campaignID, result.ID)
	assert.True(t, result.CPC.Equal(cpc))
	assert.True(t, result.AverageCPC.Equal(avgCPC))
	assert.True(t, result.MaxCPC.Equal(maxCPC))

	mockRepo.AssertExpectations(t)
}

func TestCampaignService_RefreshCampaignCPCFromGoogleAds(t *testing.T) {
	// Setup
	mockRepo := new(MockCampaignRepository)
	mockGoogleAds := new(MockGoogleAdsService)
	campaignService := services.NewCampaignService(mockRepo, mockGoogleAds)

	campaignID := uuid.New()
	googleAdsID := "123456789"
	
	campaign := &models.Campaign{
		BaseModel:           models.BaseModel{ID: campaignID},
		Name:                "Test Campaign",
		GoogleAdsCampaignID: &googleAdsID,
	}

	cpcData := &services.CPCData{
		CPC:        func() *decimal.Decimal { d := decimal.NewFromFloat(1.50); return &d }(),
		AverageCPC: func() *decimal.Decimal { d := decimal.NewFromFloat(1.25); return &d }(),
		MaxCPC:     func() *decimal.Decimal { d := decimal.NewFromFloat(2.00); return &d }(),
	}

	updatedCampaign := &models.Campaign{
		BaseModel:           models.BaseModel{ID: campaignID},
		Name:                "Test Campaign",
		GoogleAdsCampaignID: &googleAdsID,
		CPC:                 cpcData.CPC,
		AverageCPC:          cpcData.AverageCPC,
		MaxCPC:              cpcData.MaxCPC,
	}

	// Set expectations
	mockRepo.On("GetByID", campaignID).Return(campaign, nil)
	mockGoogleAds.On("GetCampaignCPC", googleAdsID).Return(cpcData, nil)
	mockRepo.On("Update", mock.MatchedBy(func(c *models.Campaign) bool {
		return c.ID == campaignID && 
			   c.CPC.Equal(*cpcData.CPC) && 
			   c.AverageCPC.Equal(*cpcData.AverageCPC) && 
			   c.MaxCPC.Equal(*cpcData.MaxCPC)
	})).Return(updatedCampaign, nil)

	// Execute
	result, err := campaignService.RefreshCampaignCPCFromGoogleAds(campaignID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, campaignID, result.ID)
	assert.True(t, result.CPC.Equal(*cpcData.CPC))
	assert.True(t, result.AverageCPC.Equal(*cpcData.AverageCPC))
	assert.True(t, result.MaxCPC.Equal(*cpcData.MaxCPC))

	mockRepo.AssertExpectations(t)
	mockGoogleAds.AssertExpectations(t)
}