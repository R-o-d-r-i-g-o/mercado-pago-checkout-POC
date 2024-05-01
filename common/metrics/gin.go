package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const metricsPath = "/metrics"

type Gin struct {
	RequestsTotal    *prometheus.CounterVec
	RequestDuration  *prometheus.HistogramVec
	RequestsInFlight *prometheus.GaugeVec
}

func NewMetricsGin() *Gin {
	m := new(Gin)

	m.RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count all http requests by status code, method and path.",
		},
		[]string{"status_code", "method", "path"},
	)

	m.RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of all HTTP requests by status code, method and path.",
		Buckets: []float64{
			0.00001, // 10µs
			0.0001,  // 100µs
			0.0002,  // 200µs
			0.0005,  // 500µs
			0.001,   // 1ms
			0.002,   // 2ms
			0.005,   // 5ms
			0.01,    // 10ms
			0.02,
			0.05,
			0.1, // 100 ms
			0.2,
			0.5,
			1.0, // 1s
			2.0,
			5.0,
		},
	}, []string{"status_code", "method", "path"})

	m.RequestsInFlight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_requests_in_progress_total",
		Help: "All the requests in progress",
	}, []string{"method"})

	return m
}

func (m *Gin) Setup(e *gin.Engine) {
	e.Use(m.Middleware())
	e.GET(metricsPath, m.Handler())
}

func (m *Gin) Handler() gin.HandlerFunc {
	h := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (m *Gin) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path

		if path == metricsPath {
			c.Next()
			return
		}

		m.RequestsInFlight.WithLabelValues(method).Inc()
		defer func() {
			m.RequestsInFlight.WithLabelValues(method).Dec()
		}()

		c.Next()

		statusCode := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9

		m.RequestDuration.WithLabelValues(statusCode, method, path).Observe(elapsed)
		m.RequestsTotal.WithLabelValues(statusCode, method, path).Inc()
	}
}
