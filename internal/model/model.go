// Package model defines the database and API request/response models for the application.
package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase opens a connection to a database, ensures the model schema is migrated
// and sets the global var `DB` to hold the open connection.
func ConnectDatabase() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("DATABASE_SERVICE")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=zumm port=%s sslmode=disable", host, username, password, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Encountered error when connecting to database")
	}

	err = database.AutoMigrate(&User{}, &Swipe{})
	if err != nil {
		return
	}

	DB = database
}
