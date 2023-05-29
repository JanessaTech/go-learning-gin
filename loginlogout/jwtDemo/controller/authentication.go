package controller

import (
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/helper"
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var auth model.AuthenticationInput
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		UserName: auth.UserName,
		Password: auth.Password,
	}
	user.BeforeSave()
	savedUser, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"user": savedUser})
}

func Login(c *gin.Context) {
	var auth model.AuthenticationInput
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := model.GetUserByName(auth.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.ValidatePassword(auth.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func AddEntry(c *gin.Context) {
	var entry model.Entry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	entry.UserId = user.ID
	user.Entries = append(user.Entries, entry)
	c.JSON(http.StatusOK, gin.H{"entry": entry})
}

func GetAllEntry(c *gin.Context) {
	user, err := helper.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"entries": user.Entries})

}
