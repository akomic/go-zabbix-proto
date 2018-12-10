go-zabbix-proto
==================

Go lang package implementing Zabbix Protocols in active mode.
Supported protocols:
- Agent sender
- Zabbix Proxy

# Tests

```bash
go test ./...
```

# Examples:

## Agent Example

```go
package main

import (
    zbxagent "go-zabbix-proto/agent"
    "log"
    "time"
)

const (
    serverHost = `localhost`
    serverPort = 10051
    agentName  = `ZabbixServer`
)

func main() {
    // New Agent
    agent := zbxagent.NewAgent(agentName, serverHost, serverPort)

    var err error

    // Sending metrics to Zabbix
    var data []*zbxagent.Metric
    data = append(data, agent.NewMetric("agent.ping", "10", time.Now().Unix()))
    data = append(data, agent.NewMetric("agent.hostname", agentName))

    var res *zbxagent.Response
    res, err = agent.Send(data)
    if err != nil {
        log.Print(err.Error())
    } else {
        log.Print("Received response: ", res.JSON)
    }
}
```

## Proxy Example

```go
package main

import (
    zbxproxy "go-zabbix-proto/proxy"
    "log"
    "time"
)

const (
    serverHost = `localhost`
    serverPort = 10051
    proxyName  = `zbxproxy-test`
)

func main() {
    // New Proxy
    proxy := zbxproxy.NewProxy(proxyName, serverHost, serverPort)

    var err error

    // Sending heartbeat to Zabbix
    var heartbeatRes *zbxproxy.ProxyResponse
    heartbeatRes, err = proxy.SendHeartbeat()
    if err != nil {
        log.Print(err.Error())
    } else {
        log.Print("Received response: ", heartbeatRes.JSON)
    }

    // Getting Proxy Config from Zabbix
    var config *zbxproxy.ProxyConfig
    config, err = proxy.GetConfig()
    if err != nil {
        log.Print(err.Error())
    } else {
        log.Println("Got", len(config.Hosts), "hosts")
    }

    // Sending host availability data
    var availabilityData []*zbxproxy.AvailabilityData
    availabilityData = append(availabilityData, &zbxproxy.AvailabilityData{
        Hostid:    10169,
        Available: 1,
    })

    var availabilityRes *zbxproxy.ProxyResponse
    availabilityRes, err = proxy.SendHostAvailability(availabilityData)

    if err != nil {
        log.Print(err.Error())
    } else {
        log.Print("Received response: ", availabilityRes.JSON)
    }

    // Sending history data
    var historyData []*zbxproxy.HistoryData
    historyData = append(historyData, &zbxproxy.HistoryData{
        Host:  `zbxagent-test`,
        Key:   `agent.ping`,
        Clock: time.Now().Unix(),
        Value: `1000`,
    })

    var historyRes *zbxproxy.ProxyResponse
    historyRes, err = proxy.SendHistory(historyData)

    if err != nil {
        log.Print(err.Error())
    } else {
        log.Print("Received response: ", historyRes.JSON)
    }
}
```
