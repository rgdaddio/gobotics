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
	if _, ok := m.devices[newDevice.Name]; ok {
		return fmt.Errorf("device already exists")
	}
	m.devices[newDevice.Name] = newDevice
	return nil
}

func (m *MockClient) UpdateDevice(device clientdevices.Device) error {
	if device, ok := m.devices[device.Name]; ok {
		m.devices[device.Name] = device
		return nil
	}
	return fmt.Errorf("device not found")
}

func (m *MockClient) FindDeviceByName(deviceName string) (clientdevices.Device, error) {
	if device, ok := m.devices[deviceName]; ok {
		return device, nil
	}
	return clientdevices.Device{}, fmt.Errorf("device not found")
}

func (m *MockClient) RemoveDeviceByName(deviceName string) error {
	if _, ok := m.devices[deviceName]; ok {
		delete(m.devices, deviceName)
		return nil
	}
	return fmt.Errorf("device not found")
}

func (m *MockClient) GetAllDevices() (clientdevices.Devices, error) {

	devices := clientdevices.Devices{}
	for _, value := range m.devices {
		devices = append(devices, value)
	}
	return devices, nil
}
