package proxy

import (
	"testing"
)

const (
	serverHost = `localhost`
	serverPort = 123
	proxyName  = `testProxy`
)

func TestSend(t *testing.T) {
	proxy := NewProxy(proxyName, serverHost, serverPort)

	_, err := proxy.SendHeartbeat()
	if err == nil {
		t.Error("Sending heartbeat should have failed")
	}
	t.Logf("error: %v", err.Error())
}
