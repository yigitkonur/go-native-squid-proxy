package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
    requestCounter = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
    )
)

func init() {
    prometheus.MustRegister(requestCounter)
}

func IncrementRequestCounter() {
    requestCounter.Inc()
}
