package utils

import (
	"AuthenticateUser/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GetUserId(ctx *gin.Context) (int, errors.CustomError) {
	tokenString, _ := ctx.Cookie("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("SecretKey")
		return []byte(secretKey), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["sub"]

	if userID, ok := userId.(float64); ok {
		return int(userID), errors.CustomError{}
	}
	return 0, errors.CustomError{Time: time.Now(), UserMessage: "please contact admin", ErrorExist: true}
}
