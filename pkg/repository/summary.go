package repository

import (
	"context"
	"gorm.io/gorm"
	"stori/cmd/logger"
	"stori/cmd/models"
)

type SummaryRepository struct {
	Database *gorm.DB
}

func NewSummaryRepository(database *gorm.DB) SummaryRepository {
	return SummaryRepository{
		Database: database,
	}
}

func (r *SummaryRepository) Save(ctx context.Context, summary *models.Summary) error {
	logger.Info(ctx, "starting...")

	response := r.Database.Create(summary)
	if response.Error != nil {
		logger.Error(ctx, response.Error.Error())

		return response.Error
	}

	return nil
}
