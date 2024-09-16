package services

import (
	"database/sql"
	"errors"

	"github.com/pavva91/bike-backend/server/internal/models"
	"github.com/pavva91/bike-backend/server/internal/repositories"
)

var Account Accounter = account{}

type Accounter interface {
	List() ([]models.Account, error)
}

type account struct{}

func (s account) List() ([]models.Account, error) {
	accounts, err := repositories.Account.List()

	if errors.Is(err, sql.ErrNoRows) {
		return []models.Account{}, nil
	}

	return accounts, err
}
