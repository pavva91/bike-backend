package repositories

import (
	"log"

	"github.com/pavva91/bike-backend/server/internal/db"
	"github.com/pavva91/bike-backend/server/internal/models"
)

var (
	Account Accounter = account{}
)

type Accounter interface {
	List() ([]models.Account, error)
	GetByID(id string) (*models.Account, error)
}

type account struct{}

func (r account) List() ([]models.Account, error) {
	rows, err := db.DB.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.Account

	for rows.Next() {
		log.Println(rows)
		var account models.Account

		err := rows.Scan(&account.ID, &account.Firstname, &account.Lastname, &account.AccountTypeID)
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

func (r account) GetByID(id string) (*models.Account, error) {
	var account models.Account
	row := db.DB.QueryRow("SELECT * FROM account WHERE id=$1", id)

	err := row.Scan(&account.ID, &account.Firstname, &account.Lastname, &account.AccountTypeID)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
