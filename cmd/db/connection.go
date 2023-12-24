package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"stori/cmd/models"
)

const DSN = "host=postgres user=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles"

func Connection() *gorm.DB {
	DB, errDB := gorm.Open(postgres.Open(
		fmt.Sprintf(DSN, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("DB_PORT"))),
		&gorm.Config{})
	if errDB != nil {
		panic(errDB)
	}

	log.Println("running migrations...")

	errMigrate := DB.AutoMigrate(&models.Summary{}, &models.Transaction{})
	if errMigrate != nil {
		panic(errMigrate)
	}

	return DB
}
