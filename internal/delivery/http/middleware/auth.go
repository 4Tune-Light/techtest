package mw

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Token from Header
		headerString := c.Request.Header.Get("Authorization")
		if headerString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		headerSplit := strings.Split(headerString, " ")
		if len(headerSplit) < 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenString := headerSplit[1]
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Get Secret Key
		secret := os.Getenv("JWT_SECRET")

		// Parse Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Validate Claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Validate Exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Validate User ID
			uid := claims["id"].(float64)
			uRole := claims["role"]
			if uid == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("x-user-id", strconv.FormatFloat(uid, 'f', -1, 64))
			c.Set("x-user-role", uRole)

			// Validate Role
			if role == "" {
				c.Next()
			} else {
				if uRole == role {
					c.Next()
				} else {
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
