package repository

import (
	"context"
	"gorm.io/gorm"
	"stori/cmd/logger"
	"stori/cmd/models"
)

type SaveDBRepository struct {
	Database *gorm.DB
}

func NewSaveDBRepository(database *gorm.DB) SaveDBRepository {
	return SaveDBRepository{
		Database: database,
	}
}

func (r *SaveDBRepository) SaveTransaction(ctx context.Context, transaction *models.Transaction) error {
	logger.Info(ctx, "starting...")

	createdTransaction := r.Database.Create(transaction)
	if createdTransaction.Error != nil {
		logger.Error(ctx, createdTransaction.Error.Error())

		return createdTransaction.Error
	}

	return nil
}

func (r *SaveDBRepository) SaveSummary(ctx context.Context, summary *models.Summary) error {
	logger.Info(ctx, "starting...")

	createdSummary := r.Database.Create(summary)
	if createdSummary.Error != nil {
		logger.Error(ctx, createdSummary.Error.Error())

		return createdSummary.Error
	}

	return nil
}
