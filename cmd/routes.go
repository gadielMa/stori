package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AppHandler struct {
	AccountHandler http.HandlerFunc
}

func (a *AppHandler) InitializeRoutes(r *chi.Mux) {
	r.Route("/stori", func(r chi.Router) {
		r.Post("/summary", a.AccountHandler)
	})
}
