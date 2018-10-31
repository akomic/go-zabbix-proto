package agent

import (
	"time"
)

// Metric structure
type Metric struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Clock int64  `json:"clock"`
}

// NewMetric constructor.
func (agent *Agent) NewMetric(key, value string, clock ...int64) *Metric {
	m := &Metric{Host: agent.Name, Key: key, Value: value}
	// use current time, if `clock` is not specified
	if m.Clock = time.Now().Unix(); len(clock) > 0 {
		m.Clock = int64(clock[0])
	}
	return m
}

// MetricPacket stucture
type MetricPacket struct {
	Request string    `json:"request"`
	Data    []*Metric `json:"data,omitempty"`
	Clock   int64     `json:"clock,omitempty"`
}
