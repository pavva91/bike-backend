package models

type AccountType struct {
	ID       uint
	Name     string
	Accounts []AccountGorm `gorm:"foreignKey:AccountTypeID"`
}
