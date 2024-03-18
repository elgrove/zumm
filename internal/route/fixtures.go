package route

import (
	"fmt"
	"strings"
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestUser(firstname, lastname string, male bool) model.User {
	location := model.UserLocation{Latitude: 51.5080, Longitude: -0.11758}
	gender := "Male"
	if male == false {
		gender = "Female"
	}
	user := model.User{
		Name:     fmt.Sprintf("%s %s", firstname, lastname),
		Age:      25,
		Gender:   gender,
		Location: location,
		Email:    strings.ToLower(fmt.Sprintf("%s@%s.com", firstname, lastname)),
		Password: "password",
	}
	return user
}

func addTestUsers(db *gorm.DB) {
	user1 := createTestUser("John", "Smith", true)
	db.Create(&user1)
	user2 := createTestUser("Jane", "Edwards", false)
	db.Create(&user2)
}

func addRandomUsers(db *gorm.DB, count int) {
	randomUsers := make([]model.User, 0, count)
	for i := 0; i < count; i++ {
		user := model.CreateRandomUser()
		randomUsers = append(randomUsers, user)
	}
	db.Create(&randomUsers)
}

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to create inmemory test db")
	}
	db.AutoMigrate(&model.User{}, &model.Swipe{})
	addTestUsers(db)
	addRandomUsers(db, 500)
	model.SetDB(db)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}

func getTokenHeaderForUser(u model.User) string {
	claims := model.UserClaims{jwt.RegisteredClaims{}, u}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	return tokenHeader
}
