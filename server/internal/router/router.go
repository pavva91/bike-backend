package router

import (
	"github.com/gorilla/mux"
	"github.com/pavva91/bike-backend/server/internal/handlers"
)

var Router *mux.Router

func NewRouter() {
	Router = mux.NewRouter()

	initializeRoutes()
}

func initializeRoutes() {
	files := Router.PathPrefix("/accounts").Subrouter()
	files.HandleFunc("", handlers.AccountsHandler.List).Methods("GET")
	files.HandleFunc("/", handlers.AccountsHandler.List).Methods("GET")
	// files.HandleFunc("/{id:[0-9]+}", handlers.AccountsHandler.Get).Methods("GET")
}
