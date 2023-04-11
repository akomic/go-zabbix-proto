package proxy

import (
	// "encoding/json"
	"fmt"

	"github.com/akomic/go-zabbix-proto/client"
)

// Proxy.
type Proxy struct {
	Name   string
	Client *client.Client
}

// Proxy constructor
func NewProxy(name string, host string, port int) (p *Proxy) {
	p = &Proxy{
		Name:   name,
		Client: client.NewClient(host, port),
	}
	return
}

// Proxy method. Sends heartbeat to Zabbix.
func (proxy *Proxy) SendHeartbeat() (response *ProxyResponse, err error) {
	packet := proxy.NewGenericPacket(`proxy heartbeat`)

	var res []byte
	res, err = proxy.Client.Send(packet)
	if err != nil {
		return
	}

	response, err = NewProxyResponse(res)
	if err != nil {
		return
	}

	if response.Response != `success` {
		err = fmt.Errorf("error sending heartbeat: %s", response.Info)
	}
	return
}

// Proxy method. Sends host availability to Zabbix.
func (proxy *Proxy) SendHostAvailability(data []*AvailabilityData) (response *ProxyResponse, err error) {
	// packet := &AvailabilityPacket{Request: `host availability`, Host: proxy.Name, Data: data}
	packet := proxy.NewAvailabilityPacket(data)

	var res []byte
	res, err = proxy.Client.Send(packet)
	if err != nil {
		return
	}

	response, err = NewProxyResponse(res)
	if err != nil {
		return
	}

	if response.Response != `success` {
		err = fmt.Errorf("error sending host availability: %s", response.Info)
	}
	return
}

// Proxy method. Sends host availability to Zabbix.
func (proxy *Proxy) SendHistory(data []*HistoryData) (response *ProxyResponse, err error) {
	// packet := &AvailabilityPacket{Request: `host availability`, Host: proxy.Name, Data: data}
	packet := proxy.NewHistoryPacket(data)

	var res []byte
	res, err = proxy.Client.Send(packet)
	if err != nil {
		return
	}

	response, err = NewProxyResponse(res)
	if err != nil {
		return
	}

	if response.Response != `success` {
		err = fmt.Errorf("error sending host availability: %s", response.Info)
	}
	return
}

// Proxy method. Send config request to Zabbix and receives configuration.
func (proxy *Proxy) GetConfig() (config *ProxyConfig, err error) {
	// packet := &Packet{Request: `proxy config`, Host: proxy.Name}
	packet := proxy.NewGenericPacket(`proxy config`)

	var res []byte
	res, err = proxy.Client.Send(packet)
	if err != nil {
		return
	}

	config, err = NewProxyConfig(res)

	return
}

// ProxyConfig
type ProxyConfig struct {
	Hosts map[float64]Host
}

// ProxyConfig contructor.
func NewProxyConfig(res []uint8) (pc *ProxyConfig, err error) {
	response, err := NewProxyConfigResponse(res)
	if err != nil {
		return
	}

	pc = &ProxyConfig{
		Hosts: response.GetHosts(),
	}
	return
}
