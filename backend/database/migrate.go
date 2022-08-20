package main

import (
	"backend/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// open database
	db, err := gorm.Open(sqlite.Open("twitter-app.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open database: twitter-app.db")
	}

	// migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Tweet{})
}
