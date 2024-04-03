package wifi

import (
	"fmt"
)

type WifiClient interface {
	GetAddresses() ([]string, error)
}

type MockWifiClient struct {
}

func (m *MockWifiClient) GetAddresses() ([]string, error) {
	return []string{"заглушка адреса 1", "заглушка адреса 2"}, nil
}

func New(client WifiClient) *WifiService {
	return &WifiService{client}
}

type WifiService struct {
	client WifiClient
}

func (w *WifiService) GetAddresses() ([]string, error) {
	return w.client.GetAddresses()
}

func main() {
	wifiService := New(&MockWifiClient{})

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		fmt.Printf("Ошибка при получении адресов: %s\n", err.Error())
		return
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
