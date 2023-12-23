package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"stori/cmd/logger"
	"stori/cmd/models"
	"stori/pkg/repository"
	"strconv"
	"strings"
)

type AccountService struct {
	MailRepository        repository.MailRepository
	TransactionRepository repository.TransactionRepository
	SummaryRepository     repository.SummaryRepository
}

func NewAccountService(mailRepository repository.MailRepository, TransactionRepository repository.TransactionRepository,
	SummaryRepository repository.SummaryRepository) AccountService {
	return AccountService{
		MailRepository:        mailRepository,
		TransactionRepository: TransactionRepository,
		SummaryRepository:     SummaryRepository,
	}
}

func (s *AccountService) ProcessSummary(ctx context.Context) error {
	logger.Info(ctx, "opening txns.csv...")

	f, errOpen := os.Open("txns.csv")
	if errOpen != nil {
		logger.Error(ctx, errOpen.Error())
		return errOpen
	}

	defer f.Close()

	logger.Info(ctx, "reading txns.csv...")

	csvReader := csv.NewReader(f)
	transactions, errRead := csvReader.ReadAll()
	if errRead != nil {
		logger.Error(ctx, errRead.Error())
		return errRead
	}

	summary, errProcess := s.process(ctx, transactions[1:])
	if errProcess != nil {
		return errProcess
	}

	errMail := s.MailRepository.Send(ctx, summary)
	if errMail != nil {
		return errMail
	}

	return nil
}

func (s *AccountService) process(ctx context.Context, transactions [][]string) (models.Summary, error) {
	var (
		monthsCount               [12]int
		creditAmount, debitAmount float64
		creditCount, debitCount   float64
		errAmount, errCount       error
	)

	logger.Info(ctx, "processing txns.csv...")

	for _, transaction := range transactions {
		creditAmount, creditCount, debitAmount, debitCount, errAmount = handleAmounts(ctx,
			transaction, creditAmount, creditCount, debitAmount, debitCount)
		if errAmount != nil {
			return models.Summary{}, errAmount
		}

		monthsCount, errCount = countMonthlyTransactions(ctx, transaction, monthsCount)
		if errCount != nil {
			return models.Summary{}, errCount
		}

		err := s.TransactionRepository.Save(ctx, &models.Transaction{
			Date:        transaction[1],
			Transaction: transaction[2],
		})
		if err != nil {
			return models.Summary{}, err
		}

	}

	summary := models.Summary{
		Balance:             big.NewFloat(0).Add(big.NewFloat(debitAmount), big.NewFloat(creditAmount)).String(),
		AvgCredit:           big.NewFloat(0).Quo(big.NewFloat(creditAmount), big.NewFloat(creditCount)).String(),
		AvgDebit:            big.NewFloat(0).Quo(big.NewFloat(debitAmount), big.NewFloat(debitCount)).String(),
		MonthlyTransactions: fmt.Sprint(monthsCount),
	}

	err := s.SummaryRepository.Save(ctx, &summary)
	if err != nil {
		return models.Summary{}, err
	}

	return summary, nil
}

func countMonthlyTransactions(ctx context.Context, transaction []string, monthsCount [12]int) ([12]int, error) {
	date := strings.Split(transaction[1], "/")
	monthIndex, err := strconv.Atoi(date[0])
	if err != nil {
		logger.Error(ctx, err.Error())
		return monthsCount, err
	}
	monthsCount[monthIndex-1] += 1

	return monthsCount, nil
}

func handleAmounts(ctx context.Context, transaction []string, creditAmount float64, creditCount float64,
	debitAmount float64, debitCount float64) (float64, float64, float64, float64, error) {
	amount, errParse := strconv.ParseFloat(strings.TrimSpace(transaction[2]), 64)
	if errParse != nil {
		logger.Error(ctx, errParse.Error())
		return 0, 0, 0, 0, errParse
	}

	if amount > 0 {
		creditAmount = creditAmount + amount
		creditCount += 1
	} else {
		debitAmount = debitAmount + amount
		debitCount += 1
	}

	return creditAmount, creditCount, debitAmount, debitCount, nil
}
