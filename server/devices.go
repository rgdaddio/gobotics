package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rgdaddio/gobotics/utils/clientdevices"
	log "github.com/sirupsen/logrus"
)

// TODO This API could be problematic when number of devices is large
// Probably need to implement paginating, and/or skip if using cassandra
func (s *Server) DevicesHandler(w http.ResponseWriter, req *http.Request) {

	var msg ServerMsg
	var statusCode int
	var devices clientdevices.Devices
	var err error

	defer func() {
		// Write response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		if statusCode == http.StatusOK {
			json.NewEncoder(w).Encode(devices)
		} else {
			json.NewEncoder(w).Encode(msg)
		}
	}()

	switch req.Method {
	case "GET":
		// List information on all devices
		devices, err = s.DeivcesClient.GetAllDevices()
		if err != nil {
			statusCode = http.StatusInternalServerError
			msg = ServerMsg{Message: "Error getting devices"}
			log.Error("Error getting devices")
			return
		}
		statusCode = http.StatusOK
	default:
		// Give an error message.
		statusCode = http.StatusMethodNotAllowed
		msg = ServerMsg{Message: "HTTP Method not supported"}
	}
}

func (s *Server) DeviceHandler(w http.ResponseWriter, req *http.Request) {

	var msg []byte
	var statusCode int

	defer func() {
		// Write response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		w.Write(msg)
	}()

	switch req.Method {
	case "GET":
		// List information on a specific device
		qmap, _ := url.ParseQuery(req.URL.RawQuery)
		ret, err := s.DeivcesClient.FindDeviceByName(qmap.Get("device"))
		if err != nil {
			statusCode = http.StatusAccepted
			msg, _ = json.Marshal(ServerMsg{Message: "Device not found"})
			return
		}
		statusCode = http.StatusOK
		msg, _ = json.Marshal(ret)
	case "POST":
		// Add a new device.
		newDevice, err := clientdevices.JsonReq2Device(req)
		if err != nil {
			statusCode = http.StatusBadRequest
			msg, _ = json.Marshal(ServerMsg{Message: fmt.Sprintf("Error decoding json: %s", err)})
			return
		}

		if newDevice.Name == "" {
			statusCode = http.StatusBadRequest
			msg, _ = json.Marshal(ServerMsg{Message: "Body missing required fields"})
			return
		}

		err = s.DeivcesClient.AddDevice(newDevice)
		if err != nil {
			if err.Error() == "device already exists" {
				statusCode = http.StatusBadRequest
				msg, _ = json.Marshal(ServerMsg{Message: "device already exists"})
				return
			}
			statusCode = http.StatusInternalServerError
			log.WithFields(log.Fields{"new_device": newDevice}).Error("Error adding new device")
			return
		}
		statusCode = http.StatusOK
		msg, _ = json.Marshal(ServerMsg{Message: "Added new device"})
	case "PUT":
		// Update an existing record.
		newDevice, err := clientdevices.JsonReq2Device(req)
		if err != nil {
			statusCode = http.StatusBadRequest
			msg, _ = json.Marshal(ServerMsg{Message: fmt.Sprintf("Error decoding json: %s", err)})
			return
		}

		if newDevice.Name == "" {
			msg, _ = json.Marshal(ServerMsg{Message: "You must specify name when trying to update device"})
			statusCode = http.StatusBadRequest
			return
		}

		err = s.DeivcesClient.UpdateDevice(newDevice)
		if err != nil {
			if err.Error() == "device not found" {
				statusCode = http.StatusNotFound
				msg, _ = json.Marshal(ServerMsg{Message: "Device not found"})
				return
			}
			statusCode = http.StatusInternalServerError
			log.WithFields(log.Fields{"new_device": newDevice}).Error("Error updating device")
			return
		}

		statusCode = http.StatusOK
		msg, _ = json.Marshal(ServerMsg{Message: "Update successful"})
	case "DELETE":
		// Remove the record.
		urlPar, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(urlPar.RawQuery)
		err := s.DeivcesClient.RemoveDeviceByName(qmap["device"][0])

		if err != nil {
			statusCode = http.StatusBadRequest
			msg, _ = json.Marshal(ServerMsg{Message: fmt.Sprintf("Couldn't Delete device: %s", qmap["device"][0])})
			return
		}
		statusCode = http.StatusOK
		msg, _ = json.Marshal(ServerMsg{Message: "Delete successful"})
	default:
		statusCode = http.StatusMethodNotAllowed
		msg, _ = json.Marshal(ServerMsg{Message: "HTTP Method not supported"})
	}
}
