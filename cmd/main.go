package main

import (
	"github.com/joho/godotenv"
	"log"
	"stori/cmd/db"
	"stori/cmd/write"
	"stori/pkg/repository"

	"github.com/go-chi/chi/v5"
	"stori/pkg/service"

	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error running application. %v", err)
	}
}

func run() error {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic(errEnv)
	}

	log.Println("starting...")

	database := db.Connection()

	mailRepository := repository.NewMailRepository()
	transactionRepository := repository.NewTransactionRepository(database)
	summaryRepository := repository.NewSummaryRepository(database)

	accountService := service.NewAccountService(mailRepository, transactionRepository, summaryRepository)
	accountHandler := write.NewAccountHandler(accountService).Handle

	app := write.AppHandler{
		AccountHandler: accountHandler,
	}

	router := chi.NewRouter()
	app.InitializeRoutes(router)

	log.Println("listening...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return err
	}

	return nil
}
