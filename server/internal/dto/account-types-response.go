package dto

import "github.com/pavva91/bike-backend/server/internal/models"

type AccountTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (dto *AccountTypeResponse) ToDTO(accountModel models.AccountType) {
	dto.ID = accountModel.ID
	dto.Name = accountModel.Name
}

func (dto *AccountTypeResponse) ToDTOs(accountsModel []models.AccountType) (userDtos []AccountTypeResponse) {
	userDtos = make([]AccountTypeResponse, len(accountsModel))
	for i, userModel := range accountsModel {
		userDtos[i].ToDTO(userModel)
	}
	return userDtos
}
