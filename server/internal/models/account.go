package models

import (
	"gorm.io/gorm"
)

type AccountGorm struct {
	gorm.Model    `swaggerignore:"true"`
	FirstName     string
	LastName      string
	AccountTypeID uint
}
