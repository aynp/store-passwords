package database

import (
	"github.com/aynp/storing-passwords/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(url string) {
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
