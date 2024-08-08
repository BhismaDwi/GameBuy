package middlewares

import (
	"GameBuy/helpers/common"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.StandardClaims
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}

		data, ok := DummyRedis[tokenString]
		if !ok {
			common.GenerateErrorResponse(c, "token invalid, please log in again")
			return
		}

		if time.Now().After(data.ExpiredAt) {
			common.GenerateErrorResponse(c, "token expired, please log in again")
			return
		}
		fmt.Println(data)
		c.Set("auth", data)
		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if common.IsEmptyField(authHeader) {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func GenerateJwtToken() (token string, err error) {
	// set token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	GenerateJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = GenerateJwtToken.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}
	return
}
