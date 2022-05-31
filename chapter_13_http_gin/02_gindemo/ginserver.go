package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) { // middleware 的作用
		s := time.Now() // 运行 handler 之前的预处理
		c.Next()        // 运行 handler
		// 运行 handler 之后的处理
		// path, response code, log latency, requestID
		reqId := 0
		requestId, _ := c.Get("RequestID")
		if m, ok := requestId.(int); ok {
			reqId = m
		}
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
			zap.Int("RequestID", reqId),
		)
	}, func(c *gin.Context) {
		c.Set("RequestID", rand.Int())
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestID, exists := c.Get("RequestID"); exists {
			h["RequestID"] = requestID
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
