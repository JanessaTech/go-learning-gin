package demomw

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FirstMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter first middleware...")
		c.Next()
		fmt.Println("exist first middleware...")
	}
}

func SecondMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter second middleware...")
		c.Next()
		fmt.Println("exist second middleware...")
	}
}

func ThirdMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter third middleware...")
		c.Next()
		fmt.Println("exist third middleware...")
	}
}

func ForthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("enter forth middleware...")
		c.Next()
		fmt.Println("exist forth middleware...")
	}
}
