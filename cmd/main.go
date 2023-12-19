package main

import (
	"github.com/joho/godotenv"
	"log"
	"stori/cmd/db"
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
	database := db.Connection()

	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic(errEnv)
	}

	log.Println("starting...")

	mailRepository := repository.NewMailRepository()
	saveDBRepository := repository.NewSaveDBRepository(database)

	accountService := service.NewAccountService(mailRepository, saveDBRepository)
	accountHandler := NewAccountHandler(accountService).Handle

	app := AppHandler{
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
