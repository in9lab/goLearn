package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

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
		req, err := http.NewRequest("GET", "http://127.0.0.1:8082/healthz", nil)
		if err != nil {
			fmt.Printf("%s", err)
		}
		lowerCaseHeader := make(http.Header)
		for k, v := range c.Request.Header {
			lowerCaseHeader[strings.ToLower((k))] = v
		}
		fmt.Printf("%s", lowerCaseHeader)
		req.Header = lowerCaseHeader
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%s", err)
		}
		if resp != nil {
			fmt.Printf("%s", resp)
		}
		glog.Infof("Client requests info: ClientIP: %s, RequestCode: %d", c.ClientIP(), http.StatusOK)
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    http.StatusOK,
		})
	})
	r.GET("/hello", rootHandler)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
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
	delay := randInt(10, 2000)
	delayTime := time.Millisecond * time.Duration(delay)
	time.Sleep(delayTime)
	c.JSON(200, gin.H{
		"message":   "hello pong",
		"code":      http.StatusOK,
		"timeSleep": delayTime,
	})
}
