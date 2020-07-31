package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/internal/package/rest"
)

// GetContext ...
func (m *DefaultMiddleware) GetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		rest.Setup(c)
	}
}
