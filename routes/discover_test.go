package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestDiscoverEndpointSuccess(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/discover", nil)
	user := createTestUser()
	claims := models.UserClaims{jwt.RegisteredClaims{}, user}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	req.Header.Add("Authorization", tokenHeader)
	router.ServeHTTP(w, req)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})
}

func TestDiscoverEndpointFailure(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/discover", nil)
	req.Header.Add("Authorization", "Bearer invalid_token_string")
	router.ServeHTTP(w, req)

	t.Run("returns 401", func(t *testing.T) {
		assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected 401")
	})
}
