package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/golang-jwt/jwt"
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

	user := models.User{
		Name:     "John Smith",
		Age:      25,
		Gender:   "Male",
		Location: "54.4,21.1",
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

func TestLoginEndpointSuccess(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()
	router := setupRouter()
	w := httptest.NewRecorder()

	userLogin := models.UserLogin{Email: "john@smith.com", Password: "heungminson7"}
	jsonData, _ := json.Marshal(userLogin)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("returns token JSON with claims", func(t *testing.T) {
		responseJSON := w.Body.Bytes()
		var tokenResponse models.TokenResponse
		json.Unmarshal(responseJSON, &tokenResponse)

		token, err := jwt.Parse(tokenResponse.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("tottenhamhotspurfootballclub"), nil
		})

		assert.NoError(t, err, "Expected token to parse without error")

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			assert.Equal(t, "john@smith.com", claims["email"], "Expected email in the claims")
			assert.Equal(t, "John Smith", claims["name"], "Expected name in the claims")
			assert.Equal(t, float64(25), claims["age"], "Expected age in the claims")
			assert.Equal(t, "Male", claims["gender"], "Expected gender in the claims")
		}
	})

}

func TestLoginEndpointFailure(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()
	router := setupRouter()
	w := httptest.NewRecorder()

	userLogin := models.UserLogin{Email: "hacker@fake.com", Password: "unauthorised"}
	jsonData, _ := json.Marshal(userLogin)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	t.Run("returns 401", func(t *testing.T) {
		assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected 401")
	})

}
