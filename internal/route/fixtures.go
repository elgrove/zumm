package route

import (
	"fmt"
	"strings"
	"zumm/internal/middleware"
	"zumm/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// createTestUser is used to create standard test users, rather than totally random.
func createTestUser(firstname, lastname string, male bool) model.User {
	location := model.UserLocation{Latitude: 51.5080, Longitude: -0.11758}
	gender := "male"
	if !male {
		gender = "female"
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

// addTestUsers is used to add a pair of standard test users to the database.
// As opposed to random users, these are static and can be used to test swipes and matches.
func addTestUsers(db *gorm.DB) {
	user1 := createTestUser("John", "Smith", true)
	db.Create(&user1)
	user2 := createTestUser("Jane", "Edwards", false)
	db.Create(&user2)
}

// addRandomUsers generates and inserts CreateRandomUser users to the database.
func addRandomUsers(db *gorm.DB, count int) {
	randomUsers := make([]model.User, 0, count)
	for i := 0; i < count; i++ {
		user := CreateRandomUser()
		randomUsers = append(randomUsers, user)
	}
	db.Create(&randomUsers)
}

// setDB is a helper function which sets the global var models.DB and can be used to patch
// the production DB with a test one.
func setDB(db *gorm.DB) {
	model.DB = db
}

// SetupTestDB returns an inmemory sqlite database pre-populated with test and random users.
// It also handles cleanup of the database after each test and subtest has completed.
func SetupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to create inmemory test db")
	}
	db.AutoMigrate(&model.User{}, &model.Swipe{})
	addTestUsers(db)
	addRandomUsers(db, 500)
	setDB(db)

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, cleanup
}

// getTokenHeaderForUser is a helper function that takes a model.User and returns
// a header string (including `Bearer `) that can be used to sign a request with JWT linked
// to the user.
func getTokenHeaderForUser(u model.User) string {
	claims := model.UserClaims{jwt.RegisteredClaims{}, u}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.JWTokenSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	return tokenHeader
}
