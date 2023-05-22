package middlewaredemo

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")

		c.Next()
		latency := time.Since(t)
		log.Println(latency)
	}
}

// http://127.0.0.1:8080/test
func AddFirstCustomMiddleWare() {
	r := gin.New()
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example")
		log.Println("example=", example)
	})
	r.Run(":8080")
}
