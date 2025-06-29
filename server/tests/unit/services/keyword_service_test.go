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

// Mock keyword repository
type MockKeywordRepository struct {
	mock.Mock
}

func (m *MockKeywordRepository) GetByAdGroupID(adGroupID uuid.UUID) ([]models.Keyword, error) {
	args := m.Called(adGroupID)
	return args.Get(0).([]models.Keyword), args.Error(1)
}

func (m *MockKeywordRepository) GetByID(keywordID uuid.UUID) (*models.Keyword, error) {
	args := m.Called(keywordID)
	return args.Get(0).(*models.Keyword), args.Error(1)
}

func (m *MockKeywordRepository) Update(keyword *models.Keyword) (*models.Keyword, error) {
	args := m.Called(keyword)
	return args.Get(0).(*models.Keyword), args.Error(1)
}

func (m *MockKeywordRepository) Create(keyword *models.Keyword) (*models.Keyword, error) {
	args := m.Called(keyword)
	return args.Get(0).(*models.Keyword), args.Error(1)
}

func (m *MockKeywordRepository) Delete(keywordID uuid.UUID) error {
	args := m.Called(keywordID)
	return args.Error(0)
}

func TestKeywordService_UpdateKeywordCPC(t *testing.T) {
	// Setup
	mockRepo := new(MockKeywordRepository)
	mockGoogleAds := new(MockGoogleAdsService)
	keywordService := services.NewKeywordService(mockRepo, mockGoogleAds)

	keywordID := uuid.New()
	cpc := decimal.NewFromFloat(1.35)
	avgCPC := decimal.NewFromFloat(1.10)
	maxCPC := decimal.NewFromFloat(1.75)

	keyword := &models.Keyword{
		BaseModel: models.BaseModel{ID: keywordID},
		Text:      "test keyword",
		MatchType: "EXACT",
	}

	updatedKeyword := &models.Keyword{
		BaseModel:  models.BaseModel{ID: keywordID},
		Text:       "test keyword",
		MatchType:  "EXACT",
		CPC:        &cpc,
		AverageCPC: &avgCPC,
		MaxCPC:     &maxCPC,
	}

	// Set expectations
	mockRepo.On("GetByID", keywordID).Return(keyword, nil)
	mockRepo.On("Update", mock.MatchedBy(func(k *models.Keyword) bool {
		return k.ID == keywordID && 
			   k.CPC.Equal(cpc) && 
			   k.AverageCPC.Equal(avgCPC) && 
			   k.MaxCPC.Equal(maxCPC)
	})).Return(updatedKeyword, nil)

	// Execute
	result, err := keywordService.UpdateKeywordCPC(keywordID, &cpc, &avgCPC, &maxCPC)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, keywordID, result.ID)
	assert.True(t, result.CPC.Equal(cpc))
	assert.True(t, result.AverageCPC.Equal(avgCPC))
	assert.True(t, result.MaxCPC.Equal(maxCPC))

	mockRepo.AssertExpectations(t)
}

func TestKeywordService_RefreshKeywordCPCFromGoogleAds(t *testing.T) {
	// Setup
	mockRepo := new(MockKeywordRepository)
	mockGoogleAds := new(MockGoogleAdsService)
	keywordService := services.NewKeywordService(mockRepo, mockGoogleAds)

	keywordID := uuid.New()
	googleAdsID := "987654321"
	
	keyword := &models.Keyword{
		BaseModel:            models.BaseModel{ID: keywordID},
		Text:                 "test keyword",
		MatchType:            "EXACT",
		GoogleAdsKeywordID:   &googleAdsID,
	}

	cpcData := &services.CPCData{
		CPC:          func() *decimal.Decimal { d := decimal.NewFromFloat(1.35); return &d }(),
		AverageCPC:   func() *decimal.Decimal { d := decimal.NewFromFloat(1.10); return &d }(),
		MaxCPC:       func() *decimal.Decimal { d := decimal.NewFromFloat(1.75); return &d }(),
		QualityScore: func() *int { i := 8; return &i }(),
		Impressions:  func() *int64 { i := int64(1000); return &i }(),
		Clicks:       func() *int64 { i := int64(50); return &i }(),
		Cost:         func() *decimal.Decimal { d := decimal.NewFromFloat(67.50); return &d }(),
	}

	updatedKeyword := &models.Keyword{
		BaseModel:            models.BaseModel{ID: keywordID},
		Text:                 "test keyword",
		MatchType:            "EXACT",
		GoogleAdsKeywordID:   &googleAdsID,
		CPC:                  cpcData.CPC,
		AverageCPC:           cpcData.AverageCPC,
		MaxCPC:               cpcData.MaxCPC,
		QualityScore:         cpcData.QualityScore,
		Impressions:          cpcData.Impressions,
		Clicks:               cpcData.Clicks,
		Cost:                 cpcData.Cost,
	}

	// Set expectations
	mockRepo.On("GetByID", keywordID).Return(keyword, nil)
	mockGoogleAds.On("GetKeywordCPC", googleAdsID).Return(cpcData, nil)
	mockRepo.On("Update", mock.MatchedBy(func(k *models.Keyword) bool {
		return k.ID == keywordID && 
			   k.CPC.Equal(*cpcData.CPC) && 
			   k.AverageCPC.Equal(*cpcData.AverageCPC) && 
			   k.MaxCPC.Equal(*cpcData.MaxCPC) &&
			   *k.QualityScore == *cpcData.QualityScore &&
			   *k.Impressions == *cpcData.Impressions &&
			   *k.Clicks == *cpcData.Clicks &&
			   k.Cost.Equal(*cpcData.Cost)
	})).Return(updatedKeyword, nil)

	// Execute
	result, err := keywordService.RefreshKeywordCPCFromGoogleAds(keywordID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, keywordID, result.ID)
	assert.True(t, result.CPC.Equal(*cpcData.CPC))
	assert.True(t, result.AverageCPC.Equal(*cpcData.AverageCPC))
	assert.True(t, result.MaxCPC.Equal(*cpcData.MaxCPC))
	assert.Equal(t, *cpcData.QualityScore, *result.QualityScore)
	assert.Equal(t, *cpcData.Impressions, *result.Impressions)
	assert.Equal(t, *cpcData.Clicks, *result.Clicks)
	assert.True(t, result.Cost.Equal(*cpcData.Cost))

	mockRepo.AssertExpectations(t)
	mockGoogleAds.AssertExpectations(t)
}