package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestDiscoverEndpointSuccess(t *testing.T) {
	_, cleanup := setupTestDB()
	defer cleanup()
	router := SetupRouter()
	w := httptest.NewRecorder()
	user := createTestUser()
	requestData := models.DiscoverRequest{Location: user.Location, DesiredGender: "Female", DesiredAgeMin: 25, DesiredAgeMax: 35}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/discover", bytes.NewBuffer(JSONData))
	claims := models.UserClaims{jwt.RegisteredClaims{}, user}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	req.Header.Add("Authorization", tokenHeader)
	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var discoverResponse models.DiscoverResponse
	parseErr := json.Unmarshal(responseJSON, &discoverResponse)
	fmt.Println("1")

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("returns results of DiscoverUserProfiles", func(t *testing.T) {
		assert.NoError(t, parseErr, "Expected response to parse without error")
	})

	t.Run("returns results sorted by nearest first", func(t *testing.T) {
		sorted := sort.SliceIsSorted(discoverResponse.Results, func(i, j int) bool {
			return discoverResponse.Results[i].DistanceFromMe < discoverResponse.Results[j].DistanceFromMe
		})
		assert.True(t, sorted, "Expected results to be sorted asc")
	})

	t.Run("returns results filtered by age as specified in posted JSON", func(t *testing.T) {
		for _, user := range discoverResponse.Results {
			if user.Age < requestData.DesiredAgeMin || user.Age > requestData.DesiredAgeMax {
				assert.FailNow(t, "Expected results to be filtered by specified age")
			}
		}
	})

	t.Run("returns results filtered by gender as specified in posted JSON", func(t *testing.T) {
		for _, user := range discoverResponse.Results {
			if user.Gender != requestData.DesiredGender {
				assert.FailNow(t, "Expected results to be filtered by specified gender")
			}
		}
	})

}

func TestDiscoverEndpointFailure(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/discover", nil)
	req.Header.Add("Authorization", "Bearer invalid_token_string")
	router.ServeHTTP(w, req)

	t.Run("returns 401", func(t *testing.T) {
		assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected 401")
	})
}
