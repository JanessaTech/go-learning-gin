package middlewaredemo

import (
	"fmt"
	demomw "hi-supergirl/go-learning-gin/MiddlewareDemo/demoMw"

	"github.com/gin-gonic/gin"
)

// http://127.0.0.1:8080/v1/api/admin/first
// http://127.0.0.1:8080/v1/api/admin/second
// http://127.0.0.1:8080/v1/api/user/third
func AddMultipleMiddlewares() {
	r := gin.New()
	r.Use(demomw.FirstMW())
	v1 := r.Group("/v1/api")
	{
		admin := v1.Group("/admin")
		{
			admin.Use() // ther mw defined in r will work
			{
				admin.GET("/first", adminFuncNoMw) // first mw is called
			}
			admin.Use(demomw.SecondMW()) // both mw defined in r and mw defined here will work
			{
				admin.GET("/second", adminFuncSecondMw) // first and second mw are called
			}

		}
		user := v1.Group("/user")
		{
			user.Use(demomw.ThirdMW())          // both mw defined in r and mw defined here will work
			user.GET("/third", userFuncThirdMw) // first and third mw are called
		}
	}
	r.Run(":8080")
}

func adminFuncNoMw(c *gin.Context) {
	fmt.Println("adminFuncNoMw ...")
	c.String(200, "adminFuncNoMw")
}

func adminFuncSecondMw(c *gin.Context) {
	fmt.Println("adminFuncSecondMW ...")
	c.String(200, "adminFuncSecondMW")
}

func userFuncThirdMw(c *gin.Context) {
	fmt.Println("userFuncThirdMw ...")
	c.String(200, "userFuncThirdMw")
}
