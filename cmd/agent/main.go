package main

import (
	"fmt"
	"time"

	"github.com/hddskull/practicus/internal/metric"
	"github.com/hddskull/practicus/internal/sender"
)

func main() {
	var m []metric.Metric
	s := sender.NewSender()
	pollCount := 0

	for {

		for i := 5; i > 0; i-- {
			pollCount++
			m = metric.GetGaugeMetrics()
			time.Sleep(2 * time.Second)
		}

		m = append(m, metric.Metric{
			Type:  "counter",
			Name:  "pollCount",
			Value: fmt.Sprint(pollCount),
		})

		for _, metric := range m {
			_, err := s.SendMetric(metric)
			if err != nil {
				panic(err)
			}
		}

	}
}
