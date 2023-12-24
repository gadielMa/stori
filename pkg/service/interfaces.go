package service

import (
	"context"
	"stori/cmd/models"
)

type SendMail interface {
	Send(ctx context.Context, summary models.Summary) error
}

type SaveSummaryRepository interface {
	Save(ctx context.Context, summary *models.Summary) error
}

type SaveTransactionRepository interface {
	Save(ctx context.Context, summary *models.Transaction) error
}
