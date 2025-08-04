// internal/config/metrics.go
package config

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics exposed to Prometheus
var (
	EventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "salesforce_events_total",
			Help: "Number of Salesforce CDC events processed",
		},
		[]string{"entity", "type"},
	)

	UpsertLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "snowflake_upsert_latency_seconds",
			Help:    "Latency of upsert operations to Snowflake",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"entity"},
	)
)

// StartMetricsServer exposes /metrics endpoint
func StartMetricsServer(port string) {
	prometheus.MustRegister(EventCounter, UpsertLatency)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("📈 Prometheus metrics listening on :%s/metrics", port)
		http.ListenAndServe(":"+port, nil)
	}()
}
