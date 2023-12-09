package models

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Migrating models...")
	err := db.AutoMigrate(
		&User{},
		&Product{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrated")
}
