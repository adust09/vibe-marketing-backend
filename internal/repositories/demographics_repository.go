package repositories

import (
	"time"

	"ads-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DemographicsRepository interface {
	Create(demographics *models.UserDemographics) error
	GetByUserID(userID uuid.UUID) (*models.UserDemographics, error)
	Update(demographics *models.UserDemographics) error
	Delete(userID uuid.UUID) error
	GetSummary() (*models.DemographicsSummary, error)
	GetPerformanceByDemographics() ([]models.DemographicsPerformance, error)
	BulkUpsert(demographics []models.UserDemographics) error
	GetStaleRecords(olderThan time.Duration) ([]models.UserDemographics, error)
}

type demographicsRepository struct {
	db *gorm.DB
}

func NewDemographicsRepository(db *gorm.DB) DemographicsRepository {
	return &demographicsRepository{db: db}
}

func (r *demographicsRepository) Create(demographics *models.UserDemographics) error {
	return r.db.Create(demographics).Error
}

func (r *demographicsRepository) GetByUserID(userID uuid.UUID) (*models.UserDemographics, error) {
	var demographics models.UserDemographics
	err := r.db.Where("user_id = ?", userID).First(&demographics).Error
	if err != nil {
		return nil, err
	}
	return &demographics, nil
}

func (r *demographicsRepository) Update(demographics *models.UserDemographics) error {
	return r.db.Save(demographics).Error
}

func (r *demographicsRepository) Delete(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserDemographics{}).Error
}

func (r *demographicsRepository) GetSummary() (*models.DemographicsSummary, error) {
	var ageDistribution []struct {
		AgeRange models.AgeRange
		Count    int
	}
	
	var genderDistribution []struct {
		Gender models.Gender
		Count  int
	}
	
	var totalUsers int64
	var lastUpdated time.Time

	if err := r.db.Model(&models.UserDemographics{}).
		Select("age_range, count(*) as count").
		Group("age_range").
		Scan(&ageDistribution).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&models.UserDemographics{}).
		Select("gender, count(*) as count").
		Group("gender").
		Scan(&genderDistribution).Error; err != nil {
		return nil, err
	}

	r.db.Model(&models.UserDemographics{}).Count(&totalUsers)
	
	r.db.Model(&models.UserDemographics{}).
		Select("MAX(last_updated_from_api)").
		Scan(&lastUpdated)

	ageMap := make(map[models.AgeRange]int)
	for _, item := range ageDistribution {
		ageMap[item.AgeRange] = item.Count
	}
	
	genderMap := make(map[models.Gender]int)
	for _, item := range genderDistribution {
		genderMap[item.Gender] = item.Count
	}

	return &models.DemographicsSummary{
		AgeDistribution:    ageMap,
		GenderDistribution: genderMap,
		TotalUsers:         int(totalUsers),
		LastUpdated:        lastUpdated,
	}, nil
}

func (r *demographicsRepository) GetPerformanceByDemographics() ([]models.DemographicsPerformance, error) {
	var performance []models.DemographicsPerformance
	
	query := `
		SELECT 
			ud.age_range,
			ud.gender,
			COUNT(ud.user_id) as user_count,
			COALESCE(AVG(c.click_through_rate), 0) as click_through_rate,
			COALESCE(AVG(c.conversion_rate), 0) as conversion_rate,
			COALESCE(AVG(c.average_cpc), 0) as average_cpc,
			COALESCE(AVG(c.roas), 0) as roas
		FROM user_demographics ud
		LEFT JOIN campaigns c ON c.user_id = ud.user_id AND c.deleted_at IS NULL
		WHERE ud.deleted_at IS NULL
		GROUP BY ud.age_range, ud.gender
		ORDER BY user_count DESC
	`
	
	err := r.db.Raw(query).Scan(&performance).Error
	return performance, err
}

func (r *demographicsRepository) BulkUpsert(demographics []models.UserDemographics) error {
	if len(demographics) == 0 {
		return nil
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, demo := range demographics {
			var existing models.UserDemographics
			err := tx.Where("user_id = ?", demo.UserID).First(&existing).Error
			
			if err == gorm.ErrRecordNotFound {
				if err := tx.Create(&demo).Error; err != nil {
					return err
				}
			} else if err != nil {
				return err
			} else {
				demo.ID = existing.ID
				demo.CreatedAt = existing.CreatedAt
				if err := tx.Save(&demo).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (r *demographicsRepository) GetStaleRecords(olderThan time.Duration) ([]models.UserDemographics, error) {
	var staleRecords []models.UserDemographics
	cutoffTime := time.Now().Add(-olderThan)
	
	err := r.db.Where("last_updated_from_api < ?", cutoffTime).Find(&staleRecords).Error
	return staleRecords, err
}