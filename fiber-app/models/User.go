package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"notNull;size:255" validate:"required,min=5,max=20"`
	Email    string `gorm:"uniqueIndex;notNull;size:255" validate:"required,email"`
	Password string `gorm:"notNull;size:255" validate:"required,min=8"`
}
