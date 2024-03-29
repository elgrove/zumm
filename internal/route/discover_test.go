package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"zumm/internal/model"

	"github.com/stretchr/testify/assert"
)

// TestDiscoverEndpointSuccess validates the behaviour of the /discover endpoint,
// when given a valid Bearer token.
// We expect that the response DiscoverUserProfiles will be sorted by nearest first as
// well as filtered by the supplied age and gender.
func TestDiscoverEndpointSuccess(t *testing.T) {
	_, cleanup := SetupTestDB()
	defer cleanup()
	router := SetupRouter()
	w := httptest.NewRecorder()
	user := createTestUser("John", "Smith", true)
	requestData := model.DiscoverRequest{Location: user.Location, DesiredGender: "female", DesiredAgeMin: 18, DesiredAgeMax: 100}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/discover", bytes.NewBuffer(JSONData))
	tokenHeader := getTokenHeaderForUser(user)
	req.Header.Add("Authorization", tokenHeader)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var discoverResponse model.DiscoverResponse
	parseErr := json.Unmarshal(responseJSON, &discoverResponse)

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

	t.Run("returns results filtered by age specified in posted JSON", func(t *testing.T) {
		for _, user := range discoverResponse.Results {
			if user.Age < requestData.DesiredAgeMin || user.Age > requestData.DesiredAgeMax {
				assert.FailNow(t, "Expected results to be filtered by specified age")
			}
		}
	})

	t.Run("returns results filtered by gender specified in posted JSON", func(t *testing.T) {
		for _, user := range discoverResponse.Results {
			if user.Gender != requestData.DesiredGender {
				assert.FailNow(t, "Expected results to be filtered by specified gender")
			}
		}
	})

}

// TestDiscoverEndpointSuccess validates the behaviour of the /discover endpoint,
// when given a invalid Bearer token.
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
