package middlewares

import (
	gin "github.com/gin-gonic/gin"
	prometheus "github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	httpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests processed",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestTotal)
}

func PrometheusCounter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		statusStr:= strconv.Itoa(ctx.Writer.Status())
		httpRequestTotal.WithLabelValues(
			ctx.Request.URL.Path,
			ctx.Request.Method,
			string(statusStr),
		).Inc()
	}
}
