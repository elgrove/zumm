package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func SetDB(db *gorm.DB) {
	DB = db
}
