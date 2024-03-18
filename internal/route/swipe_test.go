package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/internal/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func addTestSwipeMatch(db *gorm.DB) {
	janeForJohn := model.Swipe{SwiperID: 2, SwipeeID: 1, Interested: true}
	db.Create(&janeForJohn)
}

func getTestUsers(db *gorm.DB) (model.User, model.User) {
	var john model.User
	db.Where("id = ?", 1).Take(&john)
	var jane model.User
	db.Where("id = ?", 2).Take(&jane)
	return john, jane
}

func TestSwipeEndpointSuccessMatch(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()
	addTestSwipeMatch(testDB)

	router := SetupRouter()
	w := httptest.NewRecorder()

	john, jane := getTestUsers(testDB)
	requestData := model.Swipe{SwiperID: john.ID, SwipeeID: jane.ID, Interested: true}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/swipe", bytes.NewBuffer(JSONData))
	tokenHeader := getTokenHeaderForUser(john)
	req.Header.Add("Authorization", tokenHeader)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var swipeResponse model.SwipeResponse
	parseErr := json.Unmarshal(responseJSON, &swipeResponse)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("inserts swipe to db", func(t *testing.T) {
		var swipes []model.Swipe
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

	router := SetupRouter()
	w := httptest.NewRecorder()

	john, jane := getTestUsers(testDB)
	requestData := model.Swipe{SwiperID: john.ID, SwipeeID: jane.ID, Interested: false}
	JSONData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/swipe", bytes.NewBuffer(JSONData))
	tokenHeader := getTokenHeaderForUser(john)
	req.Header.Add("Authorization", tokenHeader)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	responseJSON := w.Body.Bytes()
	var swipeResponse model.SwipeResponse
	parseErr := json.Unmarshal(responseJSON, &swipeResponse)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("inserts swipe to db", func(t *testing.T) {
		var swipes []model.Swipe
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

// TODO swipe failure
