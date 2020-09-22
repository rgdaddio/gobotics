package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "github.com/rgdaddio/gobotics/server"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.HealthCheckHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, `{"message":"live"}`, strings.TrimSuffix(rr.Body.String(), "\n"))
}
