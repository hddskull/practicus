package metric

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
)

type Metric struct {
	Type  string
	Name  string
	Value string
}

// Type/Name/Value
// example: /update/gauge/otherMetric/88
func (m *Metric) GetMetricPath() string {
	return fmt.Sprintf("%s/%s/%s", m.Type, m.Name, m.Value)
}

func getGaugeMap() map[string]string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return map[string]string{
		"alloc":                          fmt.Sprint(m.Alloc),
		strings.ToLower("BuckHashSys"):   fmt.Sprint(m.BuckHashSys),
		strings.ToLower("Frees"):         fmt.Sprint(m.Frees),
		strings.ToLower("GCSys"):         fmt.Sprint(m.GCSys),
		strings.ToLower("HeapAlloc"):     fmt.Sprint(m.HeapAlloc),
		strings.ToLower("HeapIdle"):      fmt.Sprint(m.HeapIdle),
		strings.ToLower("HeapInuse"):     fmt.Sprint(m.HeapInuse),
		strings.ToLower("HeapObjects"):   fmt.Sprint(m.HeapObjects),
		strings.ToLower("HeapReleased"):  fmt.Sprint(m.HeapReleased),
		strings.ToLower("HeapSys"):       fmt.Sprint(m.HeapSys),
		strings.ToLower("LastGC"):        fmt.Sprint(m.LastGC),
		strings.ToLower("Lookups"):       fmt.Sprint(m.Lookups),
		strings.ToLower("MCacheInuse"):   fmt.Sprint(m.MCacheInuse),
		strings.ToLower("MCacheSys"):     fmt.Sprint(m.MCacheSys),
		strings.ToLower("MSpanInuse"):    fmt.Sprint(m.MSpanInuse),
		strings.ToLower("MSpanSys"):      fmt.Sprint(m.MSpanSys),
		strings.ToLower("Mallocs"):       fmt.Sprint(m.Mallocs),
		strings.ToLower("NumForcedGC"):   fmt.Sprint(uint64(m.NumForcedGC)),
		strings.ToLower("NumGC"):         fmt.Sprint(uint64(m.NumGC)),
		strings.ToLower("OtherSys"):      fmt.Sprint(m.OtherSys),
		strings.ToLower("PauseTotalNs"):  fmt.Sprint(m.PauseTotalNs),
		strings.ToLower("StackInuse"):    fmt.Sprint(m.StackInuse),
		strings.ToLower("StackSys"):      fmt.Sprint(m.StackSys),
		strings.ToLower("Sys"):           fmt.Sprint(m.Sys),
		strings.ToLower("TotalAlloc"):    fmt.Sprint(m.TotalAlloc),
		strings.ToLower("RandomValue"):   fmt.Sprint(rand.Float64()),
		strings.ToLower("GCCPUFraction"): fmt.Sprint(m.GCCPUFraction),
	}

}
func GetGaugeMetrics() []Metric {
	m := getGaugeMap()
	arr := []Metric{}
	for k, v := range m {
		arr = append(arr, Metric{
			Type:  "gauge",
			Name:  k,
			Value: v,
		})
	}
	return arr
}
