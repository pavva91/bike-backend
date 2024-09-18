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
	ListJoinAccount() ([]models.AccountType, error)
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

	var accountTypes []models.AccountType

	for rows.Next() {
		var account models.AccountType

		err := rows.Scan(&account.ID, &account.Name)
		if err != nil {
			return nil, err
		}

		accountTypes = append(accountTypes, account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accountTypes, nil
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

func (r accountType) ListJoinAccount() ([]models.AccountType, error) {
	var (
		accountTypeID        uint
		accountTypeName         string
		accountFirstname         string
		accountLastname         string
		accountID         uint
		accountType *models.AccountType
		result      []models.AccountType
	)

	rows, err := db.DB.Query("SELECT account_type.id, account_type.name, account.first_name, account.last_name, account.id FROM account_type JOIN account ON account_type.id = account.account_type_id ORDER BY account_type.id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&accountTypeID, &accountTypeName, &accountFirstname, &accountLastname, &accountID)
		if err != nil {
			return nil, err
		}
		if accountType != nil && accountType.ID == accountTypeID {
			account := &models.Account{
				Firstname: accountFirstname,
				Lastname:  accountLastname,
				ID:        accountID,
			}
			accountType.Accounts = append(accountType.Accounts, *account)
		} else {
			accountType = &models.AccountType{
				ID:   accountTypeID,
				Name: accountTypeName,
			}
			result = append(result, *accountType)
		}
	}
	return result, nil
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
