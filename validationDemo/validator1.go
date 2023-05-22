package validationdemo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func login(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err != nil {
		// there are 2 types of bind methods: Must bind and Should bind
		// For Must bind, the under hood is MustBindWith
		// For Should bind, the under hood is ShouldBindWith
		// Here the under hood is ShouldBindWith, which means user is responsible for dealing with err if binding goes wrong
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // case3
		return
	}
	if login.User != "Jane" || login.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"}) // case2
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"}) // case1
}

// http://127.0.0.1:8080/test?user=Jane&password=123    for case1
// http://127.0.0.1:8080/test?user=Jane&password=12     for case2
// http://127.0.0.1:8080/test?user=Jane&passwordd=123   for case3
func SimpleValidator() {
	r := gin.Default()
	r.GET("/test", login)
	r.Run(":8080")
}
