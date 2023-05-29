package helper

import (
	"errors"
	"fmt"
	"hi-supergirl/go-learning-gin/loginlogout/jwtDemo/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var privateKey = []byte("THIS_IS_NOT_SO_SECRET+YOU_SHOULD_DEFINITELY_CHANGE_IT")

func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func CurrentUser(ctx *gin.Context) (*model.User, error) {
	token, err := getToken(ctx)
	if err != nil {
		return &model.User{}, errors.New("failed to get token from gin.Context")
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	user, err := model.GetUserById(userId)
	if err != nil {
		fmt.Println("failed to get user by user id", userId)
	}

	return user, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""

}
