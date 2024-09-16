package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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
//	@Summary		Get
//	@Description	Get accounts
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.GetAccountResponse
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/accounts/{id} [get]

// func (h accountsHandler) Get(w http.ResponseWriter, r *http.Request) {
// 	strID := mux.Vars(r)["id"]
//
// 	i, err := strconv.Atoi(strID)
// 	if err != nil {
// 		log.Println(err)
// 		errorhandlers.BadRequestHandler(w, r, errors.New("insert valid id"))
// 		return
// 	}
// 	id := uint(i)
// 	println(id)
//
// 	// TODO: Service Get Account
// 	// TODO: DTO Account Get
// 	var res dto.GetAccountResponse
// 	res.ToDTO(accounts)
//
// 	js, err := json.Marshal(res)
// 	if err != nil {
// 		log.Println(err.Error())
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	_, err = w.Write(js)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// }
