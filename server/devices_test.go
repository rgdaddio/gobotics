package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	main "github.com/rgdaddio/gobotics/server"
	cd "github.com/rgdaddio/gobotics/utils/clientdevices"
	mock "github.com/rgdaddio/gobotics/utils/clientdevices/mock"
	"github.com/stretchr/testify/assert"
)

func TestDevicesHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		requestMethod        string
		devicesToAdd         cd.Devices
		expectedResponseCode int
		expectedResponse     string
	}{
		{
			name:                 "test when there are no devices",
			requestMethod:        "GET",
			expectedResponseCode: http.StatusOK,
			expectedResponse:     `[]`,
		},
		{
			name:          "Add one device and test get it",
			requestMethod: "GET",
			devicesToAdd: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			expectedResponseCode: http.StatusOK,
			expectedResponse:     `[{"name":"foo","platform":"test","mac_address":"::0","ip_address":"1.1.1.1"}]`,
		},
		{
			name:                 "test unsupported method",
			requestMethod:        "POST",
			expectedResponseCode: http.StatusMethodNotAllowed,
			expectedResponse:     `{"message":"HTTP Method not supported"}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			s := main.Server{}
			dc, _ := mock.NewMockClient()
			s.DeivcesClient = dc
			serverHandlerFunc := s.DevicesHandler
			endpointUri := "/client/devices"

			if len(tc.devicesToAdd) > 0 {
				for _, device := range tc.devicesToAdd {
					s.DeivcesClient.AddDevice(device)
				}
			}

			req, err := http.NewRequest(tc.requestMethod, endpointUri, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(serverHandlerFunc)
			handler.ServeHTTP(rr, req)
			status := rr.Code
			assert.Equal(t, tc.expectedResponseCode, status)
			assert.Equal(t, tc.expectedResponse, strings.TrimSuffix(rr.Body.String(), "\n"))

		})
	}
}

func TestDeviceHandlerGET(t *testing.T) {
	testCases := []struct {
		name                 string
		devicesToAdd         cd.Devices
		deviceNameToFind     string
		expectedResponseCode int
		expectedResponse     string
	}{
		{
			name:                 "no devices added nothing is returned",
			expectedResponseCode: http.StatusAccepted,
			expectedResponse:     `{"message":"Device not found"}`,
		},
		{
			name: "add one device try and find it",
			devicesToAdd: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			deviceNameToFind:     "foo",
			expectedResponseCode: http.StatusOK,
			expectedResponse:     `{"name":"foo","platform":"test","mac_address":"::0","ip_address":"1.1.1.1"}`,
		},
		{
			name: "add one device try and find a device that doesnt exist",
			devicesToAdd: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			deviceNameToFind:     "bar",
			expectedResponseCode: http.StatusAccepted,
			expectedResponse:     `{"message":"Device not found"}`,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s := main.Server{}
			dc, _ := mock.NewMockClient()
			s.DeivcesClient = dc
			serverHandlerFunc := s.DeviceHandler
			endpointUri := "/client/devices"

			if len(tc.devicesToAdd) > 0 {
				for _, device := range tc.devicesToAdd {
					s.DeivcesClient.AddDevice(device)
				}
			}

			req, err := http.NewRequest("GET", endpointUri, nil)
			if err != nil {
				t.Fatal(err)
			}

			if tc.deviceNameToFind != "" {
				q := req.URL.Query()
				q.Add("device", tc.deviceNameToFind)
				req.URL.RawQuery = q.Encode()
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(serverHandlerFunc)
			handler.ServeHTTP(rr, req)
			status := rr.Code
			assert.Equal(t, tc.expectedResponseCode, status)
			assert.Equal(t, tc.expectedResponse, strings.TrimSuffix(rr.Body.String(), "\n"))

		})
	}
}

func TestDeviceHandlerPOST(t *testing.T) {

	// no body

	// inccorect json

	// add device

	// adding same device again
	// what is the behavior here?
	// uuid?

}

func TestDeviceHandlerPUT(t *testing.T) {

}

func TestDeviceHandlerDELETE(t *testing.T) {

}
