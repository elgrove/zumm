// Package model defines the database and API request/response models for the application.
package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase opens a connection to a database, ensures the model schema is migrated
// and sets the global var `DB` to hold the open connection.
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Encountered error when connecting to database")
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return
	}

	DB = database
}
