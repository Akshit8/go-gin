// Package middleware impls middleware function for gin
package middleware

import "github.com/gin-gonic/gin"

// BasicAuth return gin basic auth middleware with permitted accounts
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"akshit": "kubernetes",
	})
}