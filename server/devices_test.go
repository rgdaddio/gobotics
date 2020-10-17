package server_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
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
		// https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
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

	s := main.Server{}
	dc, _ := mock.NewMockClient()
	s.DeivcesClient = dc
	serverHandlerFunc := s.DeviceHandler
	endpointUri := "/client/devices"

	sendReq := func(body sql.NullString) (int, string) {
		var req *http.Request
		var err error

		if !body.Valid {
			req, err = http.NewRequest("POST", endpointUri, nil)
		} else {
			req, err = http.NewRequest("POST", endpointUri, bytes.NewBuffer([]byte(body.String)))
		}

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serverHandlerFunc)
		handler.ServeHTTP(rr, req)
		return rr.Code, strings.TrimSuffix(rr.Body.String(), "\n")

	}

	// handles null body
	statusCode, resp := sendReq(sql.NullString{})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, `{"message":"Error decoding json: body is null"}`, resp)

	// no body
	statusCode, resp = sendReq(sql.NullString{String: ``, Valid: true})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, `{"message":"Error decoding json: EOF"}`, resp)

	// valid json, but none of the right fields
	statusCode, resp = sendReq(sql.NullString{String: `{"foo":"bar"}`, Valid: true})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, `{"message":"Body missing required fields"}`, resp)

	// Missing Name field
	statusCode, resp = sendReq(sql.NullString{String: `{"foo":"bar"}`, Valid: true})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, `{"message":"Body missing required fields"}`, resp)

	// add device
	statusCode, resp = sendReq(sql.NullString{String: `{"name":"foo","platform":"test","mac_address":"::0","ip_address":"1.1.1.1"}`, Valid: true})
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, `{"message":"Added new device"}`, resp)
	addedDevice, err := s.DeivcesClient.FindDeviceByName("foo")
	assert.Equal(t, nil, err)
	assert.Equal(t, addedDevice.Name, "foo")

	// adding same device again - should be denied
	sendReq(sql.NullString{String: `{"name":"bar","platform":"test","mac_address":"::0","ip_address":"1.1.1.1"}`, Valid: true})
	statusCode, resp = sendReq(sql.NullString{String: `{"name":"bar","platform":"test","mac_address":"::0","ip_address":"1.1.1.1"}`, Valid: true})
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, `{"message":"device already exists"}`, resp)
}

func TestDeviceHandlerPUT(t *testing.T) {
	testCases := []struct {
		name                 string
		requestMethod        string
		devicesToAdd         cd.Devices
		devicesToUpdate      cd.Devices
		expectedResponseCode int
		expectedResponse     string
	}{
		{
			name:          "Add no devices and try to update one",
			requestMethod: "PUT",
			devicesToUpdate: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			expectedResponseCode: http.StatusNotFound,
			expectedResponse:     `{"message":"Device not found"}`,
		},
		{
			name:          "Add one devices and try to update it",
			requestMethod: "PUT",
			devicesToAdd: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "3.3.3.3",
					Mac:      "::0",
					Platform: "test",
				},
			},
			devicesToUpdate: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			expectedResponseCode: http.StatusOK,
			expectedResponse:     `{"message":"Update successful"}`,
		},
		{
			name:          "Add one devices and try to update it",
			requestMethod: "PUT",
			devicesToAdd: []cd.Device{
				cd.Device{
					Name:     "foo",
					Ip:       "3.3.3.3",
					Mac:      "::0",
					Platform: "test",
				},
			},
			devicesToUpdate: []cd.Device{
				cd.Device{
					Name:     "bar",
					Ip:       "1.1.1.1",
					Mac:      "::0",
					Platform: "test",
				},
			},
			expectedResponseCode: http.StatusNotFound,
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

			if len(tc.devicesToUpdate) > 0 {
				for _, device := range tc.devicesToUpdate {

					deviceJson, _ := json.Marshal(device)

					req, err := http.NewRequest("PUT", endpointUri, bytes.NewBuffer([]byte(deviceJson)))
					if err != nil {
						t.Fatal(err)
					}

					rr := httptest.NewRecorder()
					handler := http.HandlerFunc(serverHandlerFunc)
					handler.ServeHTTP(rr, req)
					status := rr.Code
					assert.Equal(t, tc.expectedResponseCode, status)
					assert.Equal(t, tc.expectedResponse, strings.TrimSuffix(rr.Body.String(), "\n"))
				}
			}

		})
	}

	// PUT add device and then try updating a different device

}

func TestDeviceHandlerDELETE(t *testing.T) {
	// try delete when there are no devices

	// delete add a device and then delete

	// delete add device and then try to delete a different one
}
