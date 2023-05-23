package middlewaredemo

import (
	demomw "hi-supergirl/go-learning-gin/MiddlewareDemo/demoMw"

	"github.com/gin-gonic/gin"
)

// http:127.0.0.1:8080/v1/api/test
func GinRecoveryDemo() {
	r := gin.New()
	r.Use(demomw.FirstMW(), gin.Recovery()) // adding gin.Recovery() to prevent request chain is called more than one time when panic occurs
	r.Use(demomw.FirstMW())
	v1 := r.Group("/v1/api")
	v1.GET("/test", demomw.SetPanic())
	r.Run(":8080")
}
