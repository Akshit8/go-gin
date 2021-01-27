// Package middleware impls middleware function for gin
package middleware

import (
	"log"
	"net/http"

	"github.com/Akshit8/go-gin/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// BasicAuth return gin basic auth middleware with permitted accounts
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"akshit": "kubernetes",
	})
}

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerSchema):]

		token, err := service.NewJWTService().VerifyToken(tokenString)
		if err != nil {
			log.Printf("err verifying token: %s", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claims[Name]: ", claims["name"])
		log.Println("Claims[Admin]: ", claims["admin"])
		log.Println("Claims[Issuer]: ", claims["iss"])
		log.Println("Claims[IssuedAt]: ", claims["iat"])
		log.Println("Claims[ExpiresAt]: ", claims["exp"])
	}
}
