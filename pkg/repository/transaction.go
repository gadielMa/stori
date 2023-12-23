package repository

import (
	"context"
	"gorm.io/gorm"
	"stori/cmd/logger"
	"stori/cmd/models"
)

type TransactionRepository struct {
	Database *gorm.DB
}

func NewTransactionRepository(database *gorm.DB) TransactionRepository {
	return TransactionRepository{
		Database: database,
	}
}

func (r *TransactionRepository) Save(ctx context.Context, summary *models.Transaction) error {
	logger.Info(ctx, "starting...")

	response := r.Database.Create(summary)
	if response.Error != nil {
		logger.Error(ctx, response.Error.Error())

		return response.Error
	}

	return nil
}
