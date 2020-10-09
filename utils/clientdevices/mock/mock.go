package mock

import (
	clientdevices "github.com/rgdaddio/gobotics/utils/clientdevices"
)

//MockClient implements the ClientDevice interface
type MockClient struct {
}

func NewMockClient() (clientdevices.ClientDevices, error) {
	return nil, nil
}

func (m *MockClient) AddDevice(newDevice clientdevices.Device) error {
	return nil
}

func (m *MockClient) UpdateDevice(device clientdevices.Device) error {
	return nil
}

func (m *MockClient) FindDeviceByName(device_name string) (clientdevices.Device, error) {
	return clientdevices.Device{}, nil
}

func (m *MockClient) RemoveDeviceByName(device_name string) error {
	return nil
}

func (m *MockClient) GetAllDevices() (clientdevices.Devices, error) {
	return clientdevices.Devices{}, nil
}
