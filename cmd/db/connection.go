package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"stori/cmd/models"
)

var (
	//DB  *gorm.DB
	DSN = "host=%s user=%s password=%s dbname=%s port=%s"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connection() *gorm.DB {
	dsn := fmt.Sprintf("host=db36 port=5432 dbname=postgres user=postgres password= sslmode=disable TimeZone=Asia/Shanghai")
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE mala")
	DB = DB.Exec(createDatabaseCommand)

	log.Println("running migrations")
	DB.AutoMigrate(&models.Summary{}, &models.Transaction{})

	return DB
}
