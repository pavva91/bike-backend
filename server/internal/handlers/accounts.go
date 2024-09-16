package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pavva91/bike-backend/server/internal/dto"
	"github.com/pavva91/bike-backend/server/internal/errorhandlers"
	"github.com/pavva91/bike-backend/server/internal/services"
)

type accountsHandler struct{}

var AccountsHandler = accountsHandler{}

// List godoc
//
//	@Summary		List
//	@Description	List accounts
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.AccountResponse
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/accounts [get]
func (h accountsHandler) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := services.Account.List()
	if err != nil {
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	var re dto.AccountResponse
	res := re.ToDTOs(accounts)

	js, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// Get godoc
//
//	@Summary		Get by ID
//	@Description	Get an account by ID
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"	Format(integer)
//	@Success		200	{object}	dto.AccountResponse
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/accounts/{id} [get]
func (h accountsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	strID := mux.Vars(r)["id"]

	i, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		errorhandlers.BadRequestHandler(w, r, errors.New("insert valid id"))
		return
	}
	id := uint(i)

	account, err := services.Account.GetByID(id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			msg := fmt.Sprintf("account with id %s not found", strID)
			errorhandlers.NotFoundHandler(w, r, msg)
			return
		}

		log.Println(err.Error())
		errorhandlers.InternalServerErrorHandler(w, r)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res dto.AccountResponse
	res.ToDTO(*account)

	js, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
