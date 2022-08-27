package database

import (
	"log"
	"os"

	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Handler *gorm.DB
}

var Database DbInstance

func Connect() {
	// open database
	db, err := gorm.Open(sqlite.Open("twitter-app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to twitter-app.db")
		os.Exit(-1)
	}
	log.Println("[Database] Connected to twitter-app.db")
	log.Println("[Database] Running Migrations")

	db.AutoMigrate(&models.User{}, &models.Tweet{})

	Database = DbInstance{
		Handler: db,
	}
}
