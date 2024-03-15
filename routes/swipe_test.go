package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestSwipeEndpointSuccessMatch(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()
	addTestSwipeMatch(testDB)
	router := SetupRouter()
	w := httptest.NewRecorder()
	var john models.User
	testDB.Where("id = ?", 1).Take(&john)
	var jane models.User
	testDB.Where("id = ?", 2).Take(&jane)
	requestData := models.Swipe{SwiperID: john.ID, SwipeeID: jane.ID, Interested: true}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/swipe", bytes.NewBuffer(JSONData))
	claims := models.UserClaims{jwt.RegisteredClaims{}, john}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	req.Header.Add("Authorization", tokenHeader)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var swipeResponse models.SwipeResponse
	parseErr := json.Unmarshal(responseJSON, &swipeResponse)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("inserts swipe to db", func(t *testing.T) {
		var swipes []models.Swipe
		testDB.
			Where("swiper_id = ?", requestData.SwiperID).
			Where("swipee_id = ?", requestData.SwipeeID).
			Find(&swipes)
		assert.NotEmpty(t, swipes, "Expected swipe to be stored")
	})

	t.Run("match: response in shape expected", func(t *testing.T) {
		assert.NoError(t, parseErr)
	})

	t.Run("match: response contained match true info", func(t *testing.T) {
		assert.Equal(t, swipeResponse.Results.Matched, true)
		assert.Equal(t, *swipeResponse.Results.MatchID, jane.ID)
	})
}

func TestSwipeEndpointSuccessNoMatch(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()
	addTestSwipeNoMatch(testDB)

	router := SetupRouter()
	w := httptest.NewRecorder()
	var john models.User
	testDB.Where("id = ?", 1).Take(&john)
	var jane models.User
	testDB.Where("id = ?", 2).Take(&jane)
	requestData := models.Swipe{SwiperID: john.ID, SwipeeID: jane.ID, Interested: false}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/swipe", bytes.NewBuffer(JSONData))
	claims := models.UserClaims{jwt.RegisteredClaims{}, john}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JWTSecretKey)
	tokenHeader := fmt.Sprintf("Bearer %s", tokenString)
	req.Header.Add("Authorization", tokenHeader)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var swipeResponse models.SwipeResponse
	parseErr := json.Unmarshal(responseJSON, &swipeResponse)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("inserts swipe to db", func(t *testing.T) {
		var swipes []models.Swipe
		testDB.
			Where("swiper_id = ?", requestData.SwiperID).
			Where("swipee_id = ?", requestData.SwipeeID).
			Find(&swipes)
		assert.NotEmpty(t, swipes, "Expected swipe to be stored")
	})

	t.Run("match: response in shape expected", func(t *testing.T) {
		assert.NoError(t, parseErr)
	})

	t.Run("match: response contained match true info", func(t *testing.T) {
		assert.Equal(t, swipeResponse.Results.Matched, false)
		// TODO this isn't checking if the key is absent from the json
		assert.Nil(t, swipeResponse.Results.MatchID)
	})
}
