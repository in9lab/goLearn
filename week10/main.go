package main

import (
	"flag"
	"geek-go/metrics"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	defer glog.Flush()
	metrics.Register()
	r := gin.Default()
	r.Use(ResponseHeaders())
	r.GET("/healthz", func(c *gin.Context) {
		glog.Infof("Client requests info: ClientIP: %s, RequestCode: %d", c.ClientIP(), http.StatusOK)
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    http.StatusOK,
		})
	})
	r.GET("/hello", rootHandler)
	r.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ResponseHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for k, v := range ctx.Request.Header {
			vString := strings.Join(v, " ")
			ctx.Header(k, vString)
		}
		ctx.Header("Version", GetEnvVersion())
	}
}

func GetEnvVersion() string {
	goVersion := os.Getenv("OS")
	return goVersion
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func rootHandler(c *gin.Context) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	delay := randInt(10, 2000)
	delayTime := time.Millisecond * time.Duration(delay)
	time.Sleep(delayTime)
	c.JSON(200, gin.H{
		"message":   "hello pong",
		"code":      http.StatusOK,
		"timeSleep": delayTime,
	})
}
