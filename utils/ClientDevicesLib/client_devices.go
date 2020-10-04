package client_devices

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

// ClientDeviceLib interface to interacting with IoT Devices
type ClientDevicesLib interface {
	NewClient()
	AddDevice(newDevice Device) error
	UpdateDevice(device Device) error
	FindDeviceByName(name string) (Device, error)
	RemoveDeviceByName(name string) error
	GetAllDevices() (Devices, error)
}
