package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"notNull;size:255"`
	Email    string `gorm:"uniqueIndex;notNull;size:255"`
	Password string `gorm:"notNull;size:255"`
}
