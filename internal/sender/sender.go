package sender

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/hddskull/practicus/internal/metric"
)

const (
	URL = "http://localhost:8080/update/"
)

type Sender struct {
	client resty.Client
}

func NewSender() *Sender {
	return &Sender{
		client: *resty.New(),
	}
}

func (s *Sender) SendMetric(m metric.Metric) (*resty.Response, error) {

	fullURL := fmt.Sprint(URL, m.GetMetricPath())

	resp, err := s.client.R().
		Post(fullURL)

	if err != nil {
		return nil, err
	}

	// fmt.Println(fullURL)
	// fmt.Println(resp.StatusCode())
	// fmt.Println(resp)
	// fmt.Println()

	return resp, nil
}
