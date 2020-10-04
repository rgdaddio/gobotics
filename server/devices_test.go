package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "github.com/rgdaddio/gobotics/server"
	"github.com/stretchr/testify/assert"
)

func TestDevicesHandler(t *testing.T) {

	endpoint_uri := "/client/devices"

	req, err := http.NewRequest("GET", endpoint_uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.DevicesHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, `{"message":"live"}`, strings.TrimSuffix(rr.Body.String(), "\n"))

	// TODO Unsported method
	req, err = http.NewRequest("PUT", endpoint_uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(main.DevicesHandler)

	handler.ServeHTTP(rr, req)
	status = rr.Code
	assert.Equal(t, http.StatusMethodNotAllowed, status)
	assert.Equal(t, `{"message":"live"}`, strings.TrimSuffix(rr.Body.String(), "\n"))
}
