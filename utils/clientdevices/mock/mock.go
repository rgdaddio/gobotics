package mock

import (
	"fmt"

	clientdevices "github.com/rgdaddio/gobotics/utils/clientdevices"
)

//MockClient implements the ClientDevice interface
type MockClient struct {
	devices map[string]clientdevices.Device
}

func NewMockClient() (clientdevices.ClientDevices, error) {
	m := MockClient{}
	cdmap := make(map[string]clientdevices.Device)
	m.devices = cdmap
	return &m, nil
}

func (m *MockClient) AddDevice(newDevice clientdevices.Device) error {
	m.devices[newDevice.Name] = newDevice
	return nil
}

func (m *MockClient) UpdateDevice(device clientdevices.Device) error {
	m.devices[device.Name] = device
	return nil
}

func (m *MockClient) FindDeviceByName(device_name string) (clientdevices.Device, error) {
	if device, ok := m.devices[device_name]; ok {
		return device, nil
	}
	return clientdevices.Device{}, fmt.Errorf("Device %s not found", device_name)
}

func (m *MockClient) RemoveDeviceByName(device_name string) error {
	if _, ok := m.devices[device_name]; ok {
		delete(m.devices, device_name)
		return nil
	}
	return fmt.Errorf("Device %s not found", device_name)
}

func (m *MockClient) GetAllDevices() (clientdevices.Devices, error) {

	devices := clientdevices.Devices{}
	for _, value := range m.devices {
		devices = append(devices, value)
	}
	return devices, nil
}
