package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zumm/models"

	"github.com/stretchr/testify/assert"
)

func TestUserCreateEndpoint(t *testing.T) {
	testDB, cleanup := setupTestDB()
	defer cleanup()
	router := SetupRouter()
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
		var JohnSmith models.User
		testDB.First(&JohnSmith)
		assert.Equal(t, "john@smith.com", JohnSmith.Email, "Expected the test user John Smith to be present in db")
	})

}
