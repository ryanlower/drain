package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrainHandlerAuthentication(t *testing.T) {
	os.Setenv("AUTH_PASSWORD", "password")
	defer os.Unsetenv("AUTH_PASSWORD")

	drain := new(Drain)
	req, _ := http.NewRequest("GET", "", strings.NewReader(""))

	// Without auth in request
	w := httptest.NewRecorder()
	drain.Handler(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// With auth in request
	req.SetBasicAuth("", "password")
	w = httptest.NewRecorder()
	drain.Handler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
