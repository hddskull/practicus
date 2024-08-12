package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerUpdateMetrics(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		returnErr   bool
	}

	tests := []struct {
		name    string
		method  string
		request string
		want    want
	}{
		//check http method
		{
			name:    "Check http method: POST",
			method:  http.MethodPost,
			request: "/update/counter/someMetric/527",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check http method: Get",
			method:  http.MethodGet,
			request: "/update/counter/someMetric/527",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusUnauthorized,
				returnErr:   true,
			},
		},

		//check metric type
		{
			name:    "Check metric type: counter",
			method:  http.MethodPost,
			request: "/update/counter/someMetric/527",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check metric type: gauge",
			method:  http.MethodPost,
			request: "/update/gauge/someMetric/22.04",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check metric type: aboba",
			method:  http.MethodPost,
			request: "/update/aboba/someMetric/1337",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusBadRequest,
				returnErr:   true,
			},
		},
		//check metric name
		{
			name:    "Check metric name: someMetric",
			method:  http.MethodPost,
			request: "/update/counter/someMetric/1337",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check metric name: (empty)",
			method:  http.MethodPost,
			request: "/update/counter/",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusNotFound,
				returnErr:   true,
			},
		},
		//check metric value
		{
			name:    "Check metric value: counter/int",
			method:  http.MethodPost,
			request: "/update/counter/someMetric/527",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check metric value: counter/float",
			method:  http.MethodPost,
			request: "/update/counter/someMetric/100.05",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusBadRequest,
				returnErr:   true,
			},
		},
		{
			name:    "Check metric value: gauge/float",
			method:  http.MethodPost,
			request: "/update/gauge/otherMetric/99.05",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
		{
			name:    "Check metric value: gauge/int",
			method:  http.MethodPost,
			request: "/update/gauge/otherMetric/88",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusOK,
				returnErr:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.request, nil)
			w := httptest.NewRecorder()
			h := UpdateMetrics

			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			// assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			assert.Contains(t, result.Header.Get("Content-Type"), tt.want.contentType)
		})

	}
}
