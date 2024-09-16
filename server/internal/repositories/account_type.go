package repositories

import (
	"github.com/pavva91/bike-backend/server/internal/db"
	"github.com/pavva91/bike-backend/server/internal/models"
)

var (
	AccountType AccountTyper = accountType{}
)

type AccountTyper interface {
	List() ([]models.AccountType, error)
	GetByID(id string) (*models.AccountType, error)
}

type accountType struct{}

func (r accountType) List() ([]models.AccountType, error) {
	rows, err := db.DB.Query("SELECT * FROM account_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.AccountType

	for rows.Next() {
		// log.Println(rows)
		var account models.AccountType

		err := rows.Scan(&account.ID, &account.Name)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r accountType) GetByID(id string) (*models.AccountType, error) {
	var account models.AccountType
	row := db.DB.QueryRow("SELECT * FROM account_type WHERE id=$1", id)

	err := row.Scan(&account.ID, &account.Name)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
