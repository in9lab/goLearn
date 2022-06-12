package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	defer glog.Flush()
	r := gin.Default()
	r.Use(ResponseHeaders())
	r.GET("/healthz", func(c *gin.Context) {
		glog.Infof("Client requests info: ClientIP: %s, RequestCode: %d", c.ClientIP(), http.StatusOK)
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    http.StatusOK,
		})
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
