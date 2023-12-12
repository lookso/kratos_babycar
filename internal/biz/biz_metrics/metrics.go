package biz_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var MetricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "server",
	Subsystem: "requests",
	Name:      "duration_sec",
	Help:      "server requests duratio(sec).",
	Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1, 1.5, 2.0},
}, []string{"kind", "operation"})

var MetricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: "client",
	Subsystem: "requests",
	Name:      "code_total",
	Help:      "The total number of processed requests",
}, []string{"kind", "operation", "code", "reason"})

func Init() {
	prometheus.MustRegister(MetricSeconds, MetricRequests)
}
