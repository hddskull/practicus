package sender

import (
	"net/http"
	"testing"

	"github.com/hddskull/practicus/internal/metric"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMetrics(t *testing.T) {
	s := NewSender()

	type want struct {
		metric     metric.Metric
		statusCode int
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "send valid gauge metric",
			want: want{
				metric: metric.Metric{
					Type:  "gauge",
					Name:  "aboba",
					Value: "0.75",
				},
				statusCode: http.StatusOK,
			},
		},
		{
			name: "send valid counter metric",
			want: want{
				metric: metric.Metric{
					Type:  "counter",
					Name:  "correctMetric",
					Value: "321",
				},
				statusCode: http.StatusOK,
			},
		},
		{
			name: "send invalid counter metric",
			want: want{
				metric: metric.Metric{
					Type:  "counter",
					Name:  "incorrectMetric",
					Value: "0.75",
				},
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			require.NotNil(t, s)
			resp, err := s.SendMetric(tt.want.metric)

			require.NoError(t, err)
			assert.Equal(t, tt.want.statusCode, resp.StatusCode())

		})
	}

}
