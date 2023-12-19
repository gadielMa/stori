package models

import "gorm.io/gorm"

type Summary struct {
	gorm.Model
	Id                  uint64 `sql:"AUTO_INCREMENT"  gorm:"primary_key" json:"id"`
	Balance             string `gorm:"not null" json:"balance"`
	AvgDebit            string `gorm:"not null" json:"average_debit"`
	AvgCredit           string `gorm:"not null" json:"average_credit"`
	MonthlyTransactions string `gorm:"not null" json:"number_of_transactions"`
}
