package jwtdemo

import (
	"fmt"
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/controller"
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/middlewares"

	"github.com/gin-gonic/gin"
)

// http://127.0.0.1:8080/auth/register  POST
//
// json body
//
//	{
//	  "username" : "JanessaTech",
//	  "password" : "12345"
//	}
//
// result for example:
//
//	{
//	  "user": {
//	       "id": "85333182-5430-4323-8a16-790e2db7e4dd",
//	       "username": "JanessaTech",
//	       "password": "$2a$10$PZGmasEoiR.NAN2eQZUhL.rvBsvtaXWekklsdKdI/KcKcZ2hM7/Xu",
//	       "entries": null
//	   }
//	}
//
// http://127.0.0.1:8080/auth/login  POST
//
// json body
//
//	{
//	  "username" : "JanessaTech",
//	  "password" : "12345"
//	}
//
// result for example:
//
//	{
//	  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2ODUzNDA2NjcsImlkIjoiODUzMzMxODItNTQzMC00MzIzLThhMTYtNzkwZTJkYjdlNGRkIn0.rb9FVJSeRI2rAPgE8_kor4JLP45MMQKfa602tVRycBo"
//	}
//
// http://127.0.0.1:8080/api/addEntry   POST
//
// json body
//
//	 first post:
//		{
//		   "content" : "xxxxxxx"
//		}
//	 second post:
//		{
//		   "content" : "yyyyyyy"
//		}
//
// In postman, navigate to Authorization tab, choose "Bearer Token" in Type option, input the token string generated above
//
// result for example:
//
// for first post
//
//	{
//	   "entry": {
//	      "userID": "85333182-5430-4323-8a16-790e2db7e4dd",
//	       "content": "xxxxxxx"
//	   }
//	}
//
// for second post
//
//	{
//	   "entry": {
//	       "userID": "85333182-5430-4323-8a16-790e2db7e4dd",
//	       "content": "yyyyyyy"
//	  }
//	}
//
// http://127.0.0.1:8080/api/entries  GET
// In postman, navigate to Authorization tab, choose "Bearer Token" in Type option, input the token string generated above
//
// The result for example:
//
//	{
//	   "entries": [
//	       {
//	           "userID": "85333182-5430-4323-8a16-790e2db7e4dd",
//	           "content": "xxxxxxx"
//	       },
//	       {
//	           "userID": "85333182-5430-4323-8a16-790e2db7e4dd",
//	           "content": "yyyyyyy"
//	       }
//	  ]
//	}
func StartHttPServer() {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("register", controller.Register)
	auth.POST("login", controller.Login)

	api := r.Group("api")
	api.Use(middlewares.VerifyJwtToken())
	api.POST("/addEntry", controller.AddEntry)
	api.GET("/entries", controller.GetAllEntry)

	r.Run(":8080")
	fmt.Println("Server running on port 8000 ...")
}
