package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/app/v1/api/health/controller"
	"github.com/sofyan48/boilerplate/src/middleware"
)

type RoutesHealth struct {
	Health     controller.HealthControllerInterface
	Middleware middleware.DefaultMiddleware
}

func RoutesHealthHandler() *RoutesHealth {
	return &RoutesHealth{
		Health: controller.HealthControllerHandler(),
	}
}

type RoutesHealthInterface interface {
	Load(router *gin.Engine, version string)
}

// Load ...
func (rLoader *RoutesHealth) Load(router *gin.Engine, version string) {
	group := router.Group(version)
	group.GET("/health", rLoader.Health.HealthStatus)
	group.GET("/health/check", rLoader.Middleware.HealthChecksAuth(), rLoader.Health.HealthCheck)
}
