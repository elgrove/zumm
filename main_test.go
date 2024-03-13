package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	models.SetDB(db)

	return db, cleanup
}

func TestHelloWorldEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected 200")

	want := `{"hello":"world"}`
	assert.JSONEq(t, want, w.Body.String(), "Expected hello json")
}

func TestUserCreateEndpoint(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/create", nil)
	router.ServeHTTP(w, req)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("returns user-like JSON", func(t *testing.T) {
		responseJSON := w.Body.String()
		var user models.User
		err := json.Unmarshal([]byte(responseJSON), &user)
		if err != nil {
			t.Fatalf("Error parsing JSON: %v", err)
		}

	})

	t.Run("inserts user into DB", func(t *testing.T) {
		var count int64
		testDB.Model(&models.User{}).Count(&count)
		assert.Equal(t, int64(1), count, "Expected 1 item in the user table")
	})

}
