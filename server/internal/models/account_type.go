package models

import (
	"gorm.io/gorm"
)

type AccountTypeGorm struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Accounts   []AccountGorm `gorm:"foreignKey:AccountTypeID"`
}
