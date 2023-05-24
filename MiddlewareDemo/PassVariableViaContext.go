package middlewaredemo

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func firstMw() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCtx := context.WithValue(ctx.Request.Context(), "myname", "Janessa")
		ctx.Request = ctx.Request.WithContext(newCtx)
	}
}
func secondMw() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name, ok := ctx.Request.Context().Value("myname").(string)
		if ok {
			fmt.Println("myname = ", name)
		}

	}
}

// http://127.0.0.1:8080/test
// output:
// myname =  Janessa
func PropogateVariableViaMWDmo() {
	r := gin.New()
	r.Use(firstMw(), secondMw())
	r.GET("/test", testfunc)
	r.Run(":8080")
}

func testfunc(c *gin.Context) {
	c.String(200, "success")
}
