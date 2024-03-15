package routes

import (
	"zumm/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTestUser1() models.User {
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

func CreateTestUser2() models.User {
	location := models.UserLocation{Latitude: 51.5080, Longitude: -0.11758}
	user := models.User{
		Name:     "Jane Wilson",
		Age:      24,
		Gender:   "Female",
		Location: location,
		Email:    "jane@wilson.com",
		Password: "richarlison9",
	}
	return user
}

func addTestUsers(db *gorm.DB) {
	user1 := CreateTestUser1()
	db.Create(&user1)
	user2 := CreateTestUser2()
	db.Create(&user2)
}

func addTestSwipeMatch(db *gorm.DB) {
	janeForJohn := models.Swipe{SwiperID: 2, SwipeeID: 1, Interested: true}
	db.Create(&janeForJohn)
}

func addTestSwipeNoMatch(db *gorm.DB) {
	janeNotForJohn := models.Swipe{SwiperID: 2, SwipeeID: 1, Interested: false}
	db.Create(&janeNotForJohn)
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
	db.AutoMigrate(&models.User{}, &models.Swipe{})
	addTestUsers(db)
	addRandomUsers(db, 500)
	models.SetDB(db)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}
