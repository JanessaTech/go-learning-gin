package middlewares

import (
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyJwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helper.ValidateJWT(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "anthentication is failed"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
