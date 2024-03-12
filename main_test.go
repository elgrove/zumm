package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected 200")

	want := `{"data":"hello world"}`
	assert.Equal(t, want, w.Body.String(), "Expected hello json")
}

// TODO test for user create endpoint
func TestUserCreateEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/create", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected 200")

}
