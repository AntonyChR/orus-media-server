package middlewares

import (
	"strconv"
	"time"

	gin "github.com/gin-gonic/gin"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests processed",
		},
		[]string{"path", "method", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:"http_request_duration_miliseconds",
			Help: "Duration of http requests in miliseconds",
			Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
		},
		[]string{"path", "method", "status"},
	)

	memoryUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage",
			Help: "memory usage of process",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestTotal)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(memoryUsage)
}

func PrometheusCounter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		statusStr:= strconv.Itoa(ctx.Writer.Status())
		httpRequestTotal.WithLabelValues(
			ctx.Request.URL.Path,
			ctx.Request.Method,
			statusStr,
		).Inc()
	}
}

func PrometheusRequestDuration() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start).Milliseconds()
		statusStr:= strconv.Itoa(ctx.Writer.Status())
		requestDuration.WithLabelValues(
			ctx.Request.URL.Path,
			ctx.Request.Method,
			statusStr,
		).Observe(float64(duration))
	}
}

func PrometheusMemoryUsage() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
