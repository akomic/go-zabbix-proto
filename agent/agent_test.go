package agent

import (
	"testing"
	"time"
)

const (
	serverHost = `localhost`
	serverPort = 123
	agentName  = `testAgent`
)

func TestSend(t *testing.T) {
	agent := NewAgent(agentName, serverHost, serverPort)

	var data []*Metric
	data = append(data, agent.NewMetric("agent.ping", "10", time.Now().Unix()))
	data = append(data, agent.NewMetric("agent.hostname", agentName))

	_, err := agent.Send(data)
	if err == nil {
		t.Error("Sending metrics should have failed")
	}
	t.Logf("error: %v", err.Error())
}
