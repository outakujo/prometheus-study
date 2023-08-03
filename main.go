package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	reg := prometheus.NewRegistry()

	httpct := prometheus.NewCounterVec(prometheus.CounterOpts{Name: "http_total"},
		[]string{"path", "method"})

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		httpct,
	)

	en := gin.New()
	en.Use(gin.Recovery())
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	en.GET("metrics", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})
	en.Use(monitor(httpct))
	en.GET("tt", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})
	en.GET("a", func(c *gin.Context) {
		c.JSON(http.StatusOK, "a")
	})
	err := en.Run(":8089")
	if err != nil {
		panic(err)
	}
}

func monitor(ht *prometheus.CounterVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		ht.With(prometheus.Labels{"method": c.Request.Method,
			"path": c.Request.URL.Path}).Inc()
	}
}
