package middleware

import (
	"AuthenticateUser/domain/models"
	"AuthenticateUser/initializers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

func AuthenticateJWT(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := os.Getenv("SecretKey")
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.Users
		err = initializers.ConnectToDB().Model(&models.Users{}).First(&user, claims["sub"]).Error
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Id == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user.Id)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
