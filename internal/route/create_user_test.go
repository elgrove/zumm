package route

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/internal/model"

	"github.com/stretchr/testify/assert"
)

// TestUserCreateEndpoint validates the behaviour of the /user/create endpoint.
// When a GET request is received, a new random user is created, inserted to the DB
// and returned to the caller in JSON form.
func TestUserCreateEndpoint(t *testing.T) {
	testDB, cleanup := SetupTestDB()
	defer cleanup()
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/create", nil)
	router.ServeHTTP(w, req)

	t.Run("returns 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "Expected 200")
	})

	t.Run("returns user-like JSON", func(t *testing.T) {
		responseJSON := w.Body.Bytes()
		var user model.User
		err := json.Unmarshal(responseJSON, &user)
		if err != nil {
			t.Fatalf("Error parsing JSON: %v", err)
		}

	})

	t.Run("inserts user into DB", func(t *testing.T) {
		// TODO this does not test if a user is inserted
		var JohnSmith model.User
		testDB.First(&JohnSmith)
		assert.Equal(t, "john@smith.com", JohnSmith.Email, "Expected the test user John Smith to be present in db")
	})

}
