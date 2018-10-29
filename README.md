go-zabbix-proto
==================

Go lang package implementing Zabbix Protocols (currently only Proxy) in active mode.


# Examples:

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
        log.Print("Received response: ", heartbeatRes.Json)
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
        log.Print("Received response: ", availabilityRes.Json)
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
        log.Print("Received response: ", historyRes.Json)
    }
}
```
