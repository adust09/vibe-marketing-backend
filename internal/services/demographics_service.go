package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"ads-backend/internal/config"
	"ads-backend/internal/models"
	"ads-backend/internal/repositories"
	"github.com/google/uuid"
)

type DemographicsService interface {
	GetUserDemographics(userID uuid.UUID) (*models.UserDemographics, error)
	UpdateUserDemographics(userID uuid.UUID) error
	GetDemographicsSummary() (*models.DemographicsSummary, error)
	GetPerformanceByDemographics() ([]models.DemographicsPerformance, error)
	RefreshAllDemographics(ctx context.Context) error
	RefreshStaleRecords(ctx context.Context, olderThan time.Duration) error
}

type demographicsService struct {
	repo   repositories.DemographicsRepository
	config *config.Config
}

func NewDemographicsService(repo repositories.DemographicsRepository, config *config.Config) DemographicsService {
	return &demographicsService{
		repo:   repo,
		config: config,
	}
}

func (s *demographicsService) GetUserDemographics(userID uuid.UUID) (*models.UserDemographics, error) {
	demographics, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user demographics: %w", err)
	}
	return demographics, nil
}

func (s *demographicsService) UpdateUserDemographics(userID uuid.UUID) error {
	demographics, err := s.fetchDemographicsFromGoogleAds(userID)
	if err != nil {
		return fmt.Errorf("failed to fetch demographics from Google Ads: %w", err)
	}

	existing, err := s.repo.GetByUserID(userID)
	if err != nil {
		if err.Error() == "record not found" {
			demographics.UserID = userID
			demographics.LastUpdatedFromAPI = time.Now()
			return s.repo.Create(demographics)
		}
		return fmt.Errorf("failed to check existing demographics: %w", err)
	}

	demographics.ID = existing.ID
	demographics.UserID = userID
	demographics.CreatedAt = existing.CreatedAt
	demographics.LastUpdatedFromAPI = time.Now()
	
	return s.repo.Update(demographics)
}

func (s *demographicsService) GetDemographicsSummary() (*models.DemographicsSummary, error) {
	summary, err := s.repo.GetSummary()
	if err != nil {
		return nil, fmt.Errorf("failed to get demographics summary: %w", err)
	}
	return summary, nil
}

func (s *demographicsService) GetPerformanceByDemographics() ([]models.DemographicsPerformance, error) {
	performance, err := s.repo.GetPerformanceByDemographics()
	if err != nil {
		return nil, fmt.Errorf("failed to get performance by demographics: %w", err)
	}
	return performance, nil
}

func (s *demographicsService) RefreshAllDemographics(ctx context.Context) error {
	log.Println("Starting full demographics refresh")
	
	var allDemographics []models.UserDemographics
	
	err := s.repo.BulkUpsert(allDemographics)
	if err != nil {
		return fmt.Errorf("failed to bulk upsert demographics: %w", err)
	}
	
	log.Printf("Successfully refreshed %d demographics records", len(allDemographics))
	return nil
}

func (s *demographicsService) RefreshStaleRecords(ctx context.Context, olderThan time.Duration) error {
	log.Printf("Refreshing demographics records older than %v", olderThan)
	
	staleRecords, err := s.repo.GetStaleRecords(olderThan)
	if err != nil {
		return fmt.Errorf("failed to get stale records: %w", err)
	}
	
	if len(staleRecords) == 0 {
		log.Println("No stale demographics records found")
		return nil
	}
	
	var updatedRecords []models.UserDemographics
	for _, record := range staleRecords {
		demographics, err := s.fetchDemographicsFromGoogleAds(record.UserID)
		if err != nil {
			log.Printf("Failed to fetch demographics for user %s: %v", record.UserID, err)
			continue
		}
		
		demographics.ID = record.ID
		demographics.UserID = record.UserID
		demographics.CreatedAt = record.CreatedAt
		demographics.LastUpdatedFromAPI = time.Now()
		updatedRecords = append(updatedRecords, *demographics)
	}
	
	if len(updatedRecords) > 0 {
		err = s.repo.BulkUpsert(updatedRecords)
		if err != nil {
			return fmt.Errorf("failed to bulk upsert updated demographics: %w", err)
		}
	}
	
	log.Printf("Successfully refreshed %d out of %d stale demographics records", len(updatedRecords), len(staleRecords))
	return nil
}

func (s *demographicsService) fetchDemographicsFromGoogleAds(userID uuid.UUID) (*models.UserDemographics, error) {
	demographics := &models.UserDemographics{
		UserID:             userID,
		AgeRange:           models.AgeRangeUnknown,
		Gender:             models.GenderUnknown,
		DataSource:         "google_ads",
		PrivacyCompliant:   true,
		LastUpdatedFromAPI: time.Now(),
	}
	
	confidence := 0.5
	demographics.Confidence = &confidence
	
	return demographics, nil
}

func (s *demographicsService) mapGoogleAdsAgeRange(ageRange string) models.AgeRange {
	switch ageRange {
	case "18-24", "AGE_RANGE_18_24":
		return models.AgeRange18To24
	case "25-34", "AGE_RANGE_25_34":
		return models.AgeRange25To34
	case "35-44", "AGE_RANGE_35_44":
		return models.AgeRange35To44
	case "45-54", "AGE_RANGE_45_54":
		return models.AgeRange45To54
	case "55-64", "AGE_RANGE_55_64":
		return models.AgeRange55To64
	case "65+", "AGE_RANGE_65_PLUS":
		return models.AgeRange65Plus
	default:
		return models.AgeRangeUnknown
	}
}

func (s *demographicsService) mapGoogleAdsGender(gender string) models.Gender {
	switch gender {
	case "MALE", "male":
		return models.GenderMale
	case "FEMALE", "female":
		return models.GenderFemale
	case "OTHER", "other":
		return models.GenderOther
	default:
		return models.GenderUnknown
	}
}