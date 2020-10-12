package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "github.com/rgdaddio/gobotics/server"
	mock "github.com/rgdaddio/gobotics/utils/clientdevices/mock"
	"github.com/stretchr/testify/assert"
)

func TestDevicesHandler(t *testing.T) {

	s := main.Server{}
	dc, _ := mock.NewMockClient()
	s.DeivcesClient = dc
	serverHandlerFunc := s.DevicesHandler

	endpointUri := "/client/devices"

	req, err := http.NewRequest("GET", endpointUri, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serverHandlerFunc)

	handler.ServeHTTP(rr, req)
	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, `{"message":"live"}`, strings.TrimSuffix(rr.Body.String(), "\n"))

	// TODO Unsported method
	req, err = http.NewRequest("PUT", endpointUri, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(serverHandlerFunc)

	handler.ServeHTTP(rr, req)
	status = rr.Code
	assert.Equal(t, http.StatusMethodNotAllowed, status)
	assert.Equal(t, `{"message":"live"}`, strings.TrimSuffix(rr.Body.String(), "\n"))
}
