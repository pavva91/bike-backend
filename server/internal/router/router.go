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
	accounts := Router.PathPrefix("/accounts").Subrouter()
	accounts.HandleFunc("", handlers.AccountsHandler.List).Methods("GET")
	accounts.HandleFunc("/", handlers.AccountsHandler.List).Methods("GET")
	accounts.HandleFunc("/{id:[0-9]+}", handlers.AccountsHandler.GetByID).Methods("GET")
	accountsAccountTypes := accounts.PathPrefix("/atypes").Subrouter()
	accountsAccountTypes.HandleFunc("/{id:[0-9]+}", handlers.AccountTypesHandler.ListAccountsOfTypeByID).Methods(("GET"))

	accountTypes := Router.PathPrefix("/atypes").Subrouter()
	accountTypes.HandleFunc("", handlers.AccountTypesHandler.List).Methods("GET")
	accountTypes.HandleFunc("/", handlers.AccountTypesHandler.List).Methods("GET")
	accountTypes.HandleFunc("/{id:[0-9]+}", handlers.AccountTypesHandler.GetByID).Methods("GET")
}
