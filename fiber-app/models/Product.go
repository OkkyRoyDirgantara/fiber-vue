package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;not null;size:10"`
	Name        string `gorm:"not null;size:20"`
	Stock       int    `gorm:"not null;size:10"`
	Description string `gorm:"type:text;not null;size:255"`
	Status      bool   `gorm:"default:1"`
}
