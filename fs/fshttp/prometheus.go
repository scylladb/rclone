package fshttp

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

// Metrics provide Transport HTTP level metrics.
type Metrics struct {
	StatusCode *prometheus.CounterVec
}

func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		StatusCode: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "http",
			Name:      "status_code",
		}, []string{"host", "method", "code"}),
	}
}

var DefaultMetrics = (*Metrics)(nil)

func (m *Metrics) onResponse(req *http.Request, resp *http.Response) {
	if m == nil {
		return
	}

	var statusCode = 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	m.StatusCode.WithLabelValues(req.Host, req.Method, fmt.Sprint(statusCode)).Inc()
}
