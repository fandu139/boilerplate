package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/internal/ucase/contract"
	"github.com/sofyan48/boilerplate/src/internal/ucase/contract/health"
	"github.com/sofyan48/boilerplate/src/middleware"
)

// RoutesHealth ...
type RoutesHealth struct {
	contract   contract.Health
	middleware middleware.DefaultMiddleware
}

// RoutesHealthHandler ...
func RoutesHealthHandler() *RoutesHealth {
	return &RoutesHealth{
		contract: health.New(),
	}
}

// RoutesHealthInterface ....
type RoutesHealthInterface interface {
	Load(router *gin.Engine, version string)
}

// Load ...
func (rLoader *RoutesHealth) Load(router *gin.Engine, version string) {
	group := router.Group(version)
	group.GET("/health", rLoader.contract.HealthStatus)
	group.GET("/health/check", rLoader.middleware.HealthChecksAuth(), rLoader.contract.HealthCheck)
}
