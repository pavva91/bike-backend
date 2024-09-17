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
	GetByIDJoin(id string) (*models.AccountType, error)
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

func (r accountType) GetByIDJoin(id string) (*models.AccountType, error) {
	accountType := &models.AccountType{}

	rows, err := db.DB.Query("SELECT account_type.id, account_type.name, account.first_name, account.last_name, account.id FROM account_type JOIN account ON account_type.id = account.account_type_id WHERE account_type.id=$1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		account := &models.Account{}

		err := rows.Scan(&accountType.ID, &accountType.Name, &account.Firstname, &account.Lastname, &account.ID)
		if err != nil {
			return nil, err
		}

		accountType.Accounts = append(accountType.Accounts, *account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accountType, nil
}

func (r accountType) GetByID(id string) (*models.AccountType, error) {
	var accountType models.AccountType
	row := db.DB.QueryRow("SELECT * FROM account_type WHERE id=$1", id)

	err := row.Scan(&accountType.ID, &accountType.Name)
	if err != nil {
		return nil, err
	}
	return &accountType, nil
}
