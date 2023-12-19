package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id          uint64 `sql:"AUTO_INCREMENT"  gorm:"primary_key" json:"id"`
	Date        string `gorm:"not null" json:"date"`
	Transaction string `gorm:"not null" json:"transaction"`
}
