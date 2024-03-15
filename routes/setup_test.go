package routes

import (
	"zumm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestUser() models.User {
	user := models.User{
		Name:     "John Smith",
		Age:      25,
		Gender:   "Male",
		Location: "51.584, -0.0473",
		Email:    "john@smith.com",
		Password: "heungminson7",
	}
	return user
}

func addTestUser(db *gorm.DB) {
	user := createTestUser()
	db.Create(&user)
}

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to create inmemory test db")
	}
	db.AutoMigrate(&models.User{})
	addTestUser(db)
	models.SetDB(db)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}
