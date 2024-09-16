package dto

import "github.com/pavva91/bike-backend/server/internal/models"

type AccountResponse struct {
	ID            uint   `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	AccountTypeID uint   `json:"account-type-id" swaggerignore:"true"`
}

func (dto *AccountResponse) ToDTO(accountModel models.Account) {
	dto.ID = accountModel.ID
	dto.Firstname = accountModel.Firstname
	dto.Lastname = accountModel.Lastname
	dto.AccountTypeID = accountModel.AccountTypeID
}

func (dto *AccountResponse) ToDTOs(accountsModel []models.Account) (userDtos []AccountResponse) {
	userDtos = make([]AccountResponse, len(accountsModel))
	for i, userModel := range accountsModel {
		userDtos[i].ToDTO(userModel)
	}
	return userDtos
}
