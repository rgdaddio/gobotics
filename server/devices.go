package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

/***
    URI: /client/devices
    paths:
        GET:
            responses:
                200:
                    description: list of all devices being managed
***/
func DevicesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// List information on all devices
		devices := getClientDevices(db)
		json.NewEncoder(w).Encode(devices)

	default:
		// Give an error message.
		msg := ServerMsg{Message: "HTTP Method not supported"}
		json.NewEncoder(w).Encode(msg)
	}
}

/***
    URI: /client/die
     do a sys exit
***/
func die(w http.ResponseWriter, req *http.Request) {
	log.Printf(req.Method)
	log.Printf(req.URL.Path)
	msg := ServerMsg{Message: "killing daemon...."}
	json.NewEncoder(w).Encode(msg)
	os.Exit(1)
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
func DeviceHandler(w http.ResponseWriter, req *http.Request) {

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
		log.Println(req.RequestURI)
		url_par, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(url_par.RawQuery)
		// if no device throw 404
		ret, err := findClientDevice(db, qmap["device"][0])
		if err != nil {
			// TODO Accepted, No content, or 404?
			msg = ServerMsg{Message: "Device not found"}
			return
		}

		json.NewEncoder(w).Encode(ret)
	case "POST":
		// Add a new device.
		new_device := Device{}
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&new_device)
		addClientDevice(db, new_device)

	case "PUT":
		// Update an existing record.
		new_device := Device{}
		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&new_device)

		if (Device{}) == new_device {
			msg = ServerMsg{Message: "Device information missing in body of request"}
			statusCode = http.StatusBadRequest
			return
		}

		if new_device.Name == "" {
			msg = ServerMsg{Message: "You must specify name when trying to update device"}
			statusCode = http.StatusBadRequest
			return
		}

		if new_device.Platform == "" && new_device.Mac == "" && new_device.Ip == "" {
			msg = ServerMsg{Message: "Please give information for update"}
		}

		updateClientDevice(db, new_device)

	case "DELETE":
		// Remove the record.
		url_par, _ := url.Parse(req.RequestURI)
		qmap, _ := url.ParseQuery(url_par.RawQuery)
		ret := remove_client_device(db, qmap["device"][0])
		if ret > 0 {
			msg = ServerMsg{Message: "Device Removed"}
		} else {
			msg = ServerMsg{Message: "Device name not found"}
		}
	default:
		msg = ServerMsg{Message: "HTTP Method not supported"}
		statusCode = http.StatusMethodNotAllowed
	}
}
