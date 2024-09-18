package services

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/pavva91/bike-backend/server/internal/models"
	"github.com/pavva91/bike-backend/server/internal/repositories"
)

var AccountType AccountTyper = accountType{}

type AccountTyper interface {
	List() ([]models.AccountType, error)
	GetByIDJoin(id uint) (*models.AccountType, error)
}

type accountType struct{}

func (s accountType) List() ([]models.AccountType, error) {
	accountTypes, err := repositories.AccountType.ListJoinAccount()

	if errors.Is(err, sql.ErrNoRows) {
		return []models.AccountType{}, nil
	}

	return accountTypes, err
}

func (s accountType) GetByIDJoin(id uint) (*models.AccountType, error) {
	strID := strconv.Itoa(int(id))
	accountType, err := repositories.AccountType.GetByIDJoin(strID)

	return accountType, err
}
