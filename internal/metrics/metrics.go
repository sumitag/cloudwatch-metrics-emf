package metrics

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Metrics struct {
	Timestamp time.Time `json:"timestamp"`
	Metrics   []Metric  `json:"metrics"`
	Metadata  Metadata  `json:"metadata"`
}

type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type Metadata struct {
	ServiceName string `json:"service_name"`
	ServiceType string `json:"service_type"`
}

func CollectMetrics() Metrics {
	// Collect CPU and memory usage
	cpuUsage, _ := cpu.Percent(0, false)
	memStats, _ := mem.VirtualMemory()

	// Collect network I/O
	netStats, _ := net.IOCounters(false)

	return Metrics{
		Timestamp: time.Now(),
		Metrics: []Metric{
			{
				Name:  "CPUUtilization",
				Value: cpuUsage[0],
				Unit:  "Percent",
			},
			{
				Name:  "MemoryUsage",
				Value: memStats.UsedPercent,
				Unit:  "Percent",
			},
			{
				Name:  "BytesSent",
				Value: float64(netStats[0].BytesSent),
				Unit:  "Bytes",
			},
			{
				Name:  "BytesReceived",
				Value: float64(netStats[0].BytesRecv),
				Unit:  "Bytes",
			},
		},
		Metadata: Metadata{
			ServiceName: "MyService",
			ServiceType: "MyServiceType",
		},
	}
}

// ToEMF converts Metrics to EMF JSON string
func (m Metrics) ToEMF() (string, error) {
	emfData, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(emfData), nil
}
