package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rgdaddio/gobotics/utils/clientdevices"
	log "github.com/sirupsen/logrus"
)

/***
    URI: /client/devices
    paths:
        GET:
            responses:
                200:
                    description: list of all devices being managed
***/
func (s *Server) DevicesHandler(w http.ResponseWriter, req *http.Request) {

	var msg ServerMsg
	var statusCode int

	defer func() {
		// Write response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(msg)
	}()

	switch req.Method {
	case "GET":
		// List information on all devices
		devices, err := s.DeivcesClient.GetAllDevices()
		if err != nil {
			statusCode = http.StatusInternalServerError
			msg = ServerMsg{Message: "Error getting devices"}
			log.Error("Error getting devices")
			return
		}
		json.NewEncoder(w).Encode(devices)

	default:
		// Give an error message.
		msg := ServerMsg{Message: "HTTP Method not supported"}
		json.NewEncoder(w).Encode(msg)
	}
}

/***
    URI: /client/device
    paths:
        GET:
            query parameters:
                device: name of device to get info for
            responses:
                200:
                    description: return information on device
	POST:
            parameters:
                name string : name of device
                mac_address string : mac_address of device
                ip_address  string: ip of device
                platform string: platform
            responses:
                200:
                    description: device entry created
	PUT:
            parameters:
                name string : name of device
                mac_address string : mac_address of device
                ip_address  string: ip of device
                platform string: platform
            responses:
                200:
                    description: device entry updated based on information in body
	DELETE:
            query parameters:
                device: name of device to get info for
            responses:
                200:
                    description: device removed
***/
func (s *Server) DeviceHandler(w http.ResponseWriter, req *http.Request) {

	var msg ServerMsg
	var statusCode int

	defer func() {
		// Write response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(msg)
	}()

	switch req.Method {
	case "GET":
		// List information on a specific device
		urlPar, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(urlPar.RawQuery)
		ret, err := s.DeivcesClient.FindDeviceByName(qmap["device"][0])
		if err != nil {
			statusCode = http.StatusAccepted
			msg = ServerMsg{Message: "Device not found"}
			return
		}
		statusCode = http.StatusOK
		json.NewEncoder(w).Encode(ret)
	case "POST":
		// Add a new device.
		newDevice, err := clientdevices.JsonReq2Device(req)

		if err != nil {
			statusCode = http.StatusBadRequest
			msg = ServerMsg{Message: fmt.Sprintf("Error decoding json: %s", err)}
			return
		}
		err = s.DeivcesClient.AddDevice(newDevice)
		if err != nil {
			statusCode = http.StatusInternalServerError
			log.WithFields(log.Fields{"new_device": newDevice}).Error("Error adding new device")
			return
		}
		statusCode = http.StatusOK
		msg = ServerMsg{Message: "Adding new device successful"}
	case "PUT":
		// Update an existing record.
		newDevice, err := clientdevices.JsonReq2Device(req)
		if err != nil {
			statusCode = http.StatusBadRequest
			msg = ServerMsg{Message: fmt.Sprintf("Error decoding json: %s", err)}
			return
		}

		if newDevice.Name == "" {
			msg = ServerMsg{Message: "You must specify name when trying to update device"}
			statusCode = http.StatusBadRequest
			return
		}
		err = s.DeivcesClient.UpdateDevice(newDevice)
		if err != nil {
			statusCode = http.StatusInternalServerError
			log.WithFields(log.Fields{"new_device": newDevice}).Error("Error updating device")
			return
		}
		statusCode = http.StatusOK
		msg = ServerMsg{Message: "Update successful"}
	case "DELETE":
		// Remove the record.
		urlPar, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(urlPar.RawQuery)
		err := s.DeivcesClient.RemoveDeviceByName(qmap["device"][0])

		if err != nil {
			statusCode = http.StatusBadRequest
			msg = ServerMsg{Message: fmt.Sprintf("Couldn't Delete device: %s", qmap["device"][0])}
			return
		}
		statusCode = http.StatusOK
		msg = ServerMsg{Message: "Delete successful"}
	default:
		msg = ServerMsg{Message: "HTTP Method not supported"}
		statusCode = http.StatusMethodNotAllowed
	}
}
