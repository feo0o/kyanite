package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_counter",
		Help: "HTTP(S) requests count",
	}, []string{
		"method", "path",
	})

	httpRquestTime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_request_use_time",
		Help: "HTTP(S) requests use time",
	}, []string{
		"method", "path",
	})

	httpRquestStatus = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_status",
		Help: "HTTP(S) requests status count",
	}, []string{
		"method", "path", "status",
	})
)

func init() {
	prometheus.MustRegister(httpRequestCounter, httpRquestTime, httpRquestStatus)
}

func PrometheusMetrics(c *gin.Context) {
	startTime := time.Now()
	httpRequestCounter.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Inc()
	c.Next()
	endTime := time.Now()
	httpRquestTime.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Set(float64(endTime.Sub(startTime)))
	httpRquestStatus.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath(), "status": strconv.Itoa(c.Writer.Status())}).Inc()
}
