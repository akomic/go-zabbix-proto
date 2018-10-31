package agent

import (
	"encoding/json"
	"fmt"
	"go-zabbix-proto/client"
	"time"
)

// Agent stucture
type Agent struct {
	Name   string
	Client *client.Client
}

// NewAgent constructor
func NewAgent(name string, host string, port int) (agent *Agent) {
	agent = &Agent{
		Name:   name,
		Client: client.NewClient(host, port),
	}
	return
}

// NewMetricPacket constructor.
func (agent *Agent) NewMetricPacket(data []*Metric, clock ...int64) *client.Packet {
	mp := &MetricPacket{Request: `sender data`, Data: data}
	// use current time, if `clock` is not specified
	if mp.Clock = time.Now().Unix(); len(clock) > 0 {
		mp.Clock = int64(clock[0])
	}
	jsonData, _ := json.Marshal(mp)
	packet := &client.Packet{Request: `sender data`, Data: jsonData}
	return packet
}

// Send method. Sends metrics to Zabbix.
func (agent *Agent) Send(data []*Metric, clock ...int64) (response *Response, err error) {
	packet := agent.NewMetricPacket(data)

	var res []byte
	res, err = agent.Client.Send(packet)
	if err != nil {
		return
	}

	response, err = NewResponse(res)
	if err != nil {
		return
	}

	if response.Response != `success` {
		err = fmt.Errorf("Error sending: %s", response.Info)
	}
	return
}
