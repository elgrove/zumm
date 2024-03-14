package routes

import (
	"zumm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	user := models.User{
		Name:     "John Smith",
		Age:      25,
		Gender:   "Male",
		Location: "51.584, -0.0473",
		Email:    "john@smith.com",
		Password: "heungminson7",
	}
	db.Create(&user)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	models.SetDB(db)

	return db, cleanup
}
