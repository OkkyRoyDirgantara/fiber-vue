package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;not null;size:10" validate:"required,min=3,max=10"`
	Name        string `gorm:"not null;size:20" validate:"required,min=3,max=20"`
	Stock       int    `gorm:"not null;size:10" validate:"required,min=1"`
	Description string `gorm:"type:text;not null;size:255" validate:"required,min=3,max=255"`
	Status      bool   `gorm:"default:1"`
}
