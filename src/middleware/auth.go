package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Auth ...
func (m *DefaultMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

// HealthChecksAuth ...
func (m *DefaultMiddleware) HealthChecksAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("APP_USERNAME"): os.Getenv("APP_PASSWORD"),
	})
}
