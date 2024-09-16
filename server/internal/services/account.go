package services

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/pavva91/bike-backend/server/internal/models"
	"github.com/pavva91/bike-backend/server/internal/repositories"
)

var Account Accounter = account{}

type Accounter interface {
	List() ([]models.Account, error)
	GetByID(id uint) (*models.Account, error)
}

type account struct{}

func (s account) List() ([]models.Account, error) {
	accounts, err := repositories.Account.List()

	if errors.Is(err, sql.ErrNoRows) {
		return []models.Account{}, nil
	}

	return accounts, err
}

func (s account) GetByID(id uint) (*models.Account, error) {
	strID := strconv.Itoa(int(id))
	account, err := repositories.Account.GetByID(strID)

	return account, err
}
