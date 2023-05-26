package sessiondemo

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

var secret = []byte("secret")

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("AuthRequired", "session Id=", session.ID())
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// http://127.0.0.1:8080/login
// http://127.0.0.1:8080/logout
// http://127.0.0.1:8080/private/me
// http://127.0.0.1:8080/private/status
func LoginLogoutWithSessionDemo() {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))
	r.POST("/login", login)
	r.GET("/logout", logout)

	// Private group, require authentication to access
	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}
	r.Use(gin.Logger())
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}

}
func login(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("login", "session Id=", session.ID())
	session.Set(userkey, "JanessaTech")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})

}
func logout(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("logout", "session Id=", session.ID())
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("me", "session Id=", session.ID())
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// status is the handler that will tell the user whether it is logged in or not.
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
