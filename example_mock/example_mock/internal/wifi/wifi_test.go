package wifi_test

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type WiFiService struct {
	WiFi WiFi
}

func New(wifi WiFi) *WiFiService {
	return &WiFiService{WiFi: wifi}
}

func (service WiFiService) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}
	var addrs []net.HardwareAddr

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service WiFiService) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}
	var nameList []string

	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}
	return nameList, nil
}

type MockWiFi struct {
}

func (m *MockWiFi) Interfaces() ([]*wifi.Interface, error) {
	return []*wifi.Interface{
		&wifi.Interface{
			HardwareAddr: net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
			Name:         "wifi0",
		},
		&wifi.Interface{
			HardwareAddr: net.HardwareAddr{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB},
			Name:         "wifi1",
		},
	}, nil
}

func TestWiFiService_GetAddresses(t *testing.T) {
	mockWiFi := &MockWiFi{}
	wifiService := New(mockWiFi)

	addrs, err := wifiService.GetAddresses()

	expectedAddrs := []net.HardwareAddr{
		net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
		net.HardwareAddr{0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB},
	}

	assert.NoError(t, err, "ожидается отсутствие ошибок")
	assert.Equal(t, expectedAddrs, addrs, "ожидается конкретные адреса")
}

func TestWiFiService_GetNames(t *testing.T) {
	mockWiFi := &MockWiFi{}
	wifiService := New(mockWiFi)

	names, err := wifiService.GetNames()

	expectedNames := []string{"wifi0", "wifi1"}

	assert.NoError(t, err, "ожидается отсутствие ошибок")
	assert.Equal(t, expectedNames, names, "ожидается конкретные имена интерфейсов")
}
