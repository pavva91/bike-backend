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

type accountTypesHandler struct{}

var AccountTypesHandler = accountTypesHandler{}

// List godoc
//
//	@Summary		List
//	@Description	List accounts
//	@Tags			AccountTypes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.AccountTypeResponse
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/atypes [get]
func (h accountTypesHandler) List(w http.ResponseWriter, r *http.Request) {
	accountTypes, err := services.AccountType.List()
	if err != nil {
		errorhandlers.BadRequestHandler(w, r, err)
		return
	}

	var re dto.AccountTypeResponse
	res := re.ToDTOs(accountTypes)

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
//	@Description	Get an account type by ID
//	@Tags			AccountTypes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"AccountType ID"	Format(integer)
//	@Success		200	{object}	dto.AccountTypeResponse
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/atypes/{id} [get]
func (h accountTypesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	strID := mux.Vars(r)["id"]

	i, err := strconv.Atoi(strID)
	if err != nil {
		log.Println(err)
		errorhandlers.BadRequestHandler(w, r, errors.New("insert valid id"))
		return
	}
	id := uint(i)

	accountType, err := services.AccountType.GetByID(id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			msg := fmt.Sprintf("account type with id %s not found", strID)
			errorhandlers.NotFoundHandler(w, r, msg)
			return
		}

		log.Println(err.Error())
		errorhandlers.InternalServerErrorHandler(w, r)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res dto.AccountTypeResponse
	res.ToDTO(*accountType)

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
