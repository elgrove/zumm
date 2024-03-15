package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestLoginEndpointSuccess(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()
	router := SetupRouter()
	w := httptest.NewRecorder()

	user := CreateTestUser1()
	userLogin := models.LoginRequest{Email: user.Email, Password: user.Password}
	jsonData, _ := json.Marshal(userLogin)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("returns token JSON with claims", func(t *testing.T) {
		responseJSON := w.Body.Bytes()
		var tokenResponse models.LoginResponse
		json.Unmarshal(responseJSON, &tokenResponse)
		token, err := jwt.Parse(tokenResponse.Token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("tottenhamhotspurfootballclub"), nil
		})
		assert.NoError(t, err, "Expected token to parse without error")
		if claims, ok := token.Claims.(models.UserClaims); ok && token.Valid {
			assert.Equal(t, user.Email, claims.Email, "Expected email in the claims")
			assert.Equal(t, user.Name, claims.Name, "Expected name in the claims")
			assert.Equal(t, user.Age, claims.Age, "Expected age in the claims")
			assert.Equal(t, user.Gender, claims.Gender, "Expected gender in the claims")
		}
	})
}

func TestLoginEndpointFailure(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()
	router := SetupRouter()
	w := httptest.NewRecorder()

	userLogin := models.LoginRequest{Email: "hacker@fake.com", Password: "unauthorised"}
	jsonData, _ := json.Marshal(userLogin)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	t.Run("returns 401", func(t *testing.T) {
		assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected 401")
	})

}
