package clientdevices

import (
	"encoding/json"
	"net/http"
)

// Device represents a single IoT Device or Asset in the field
type Device struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Mac      string `json:"mac_address"`
	Ip       string `json:"ip_address"`
	//  Uptime time.Time `json:"uptime"`
}

// Devices - a list of multiple devices
type Devices []Device

// ClientDevice interface to interacting with IoT Devices
type ClientDevices interface {
	AddDevice(newDevice Device) error
	UpdateDevice(device Device) error
	FindDeviceByName(name string) (Device, error)
	RemoveDeviceByName(name string) error
	GetAllDevices() (Devices, error)
}

func JsonReq2Device(req *http.Request) (Device, error) {
	device := Device{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&device)
	return device, err
}
