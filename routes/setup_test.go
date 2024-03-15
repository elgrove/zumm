package routes

import (
	"zumm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestUser() models.User {
	location := models.UserLocation{Latitude: 51.5080, Longitude: -0.11758}
	user := models.User{
		Name:     "John Smith",
		Age:      25,
		Gender:   "Male",
		Location: location,
		Email:    "john@smith.com",
		Password: "heungminson7",
	}
	return user
}

func addTestUser(db *gorm.DB) {
	user := createTestUser()
	db.Create(&user)
}

func addRandomUsers(db *gorm.DB, count int) {
	randomUsers := make([]models.User, 0, count)
	for i := 0; i < count; i++ {
		user := models.CreateRandomUser()
		randomUsers = append(randomUsers, user)
	}
	db.Create(&randomUsers)
}

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to create inmemory test db")
	}
	db.AutoMigrate(&models.User{})
	addTestUser(db)
	addRandomUsers(db, 500)
	models.SetDB(db)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}
