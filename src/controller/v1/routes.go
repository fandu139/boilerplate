package users

import (
	"github.com/gin-gonic/gin"
	health "github.com/ralali/event-api/src/controller/v1/health"
	login "github.com/ralali/event-api/src/controller/v1/login"
	sso "github.com/ralali/event-api/src/controller/v1/sso"
	users "github.com/ralali/event-api/src/controller/v1/users"

	healthService "github.com/ralali/event-api/src/service/v1/health"
	loginServices "github.com/ralali/event-api/src/service/v1/login"
	ssoServices "github.com/ralali/event-api/src/service/v1/sso"
	"github.com/ralali/event-api/src/util/middleware"
)

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
}

// V1Router Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Router(router *gin.Engine) {
	
	// Health Handler Routes
	healthHandler := &health.V1HealthController{
		HealthService: healthService.V1HealthCheckHandler(),
	}

	//********* Calling Handler To Routers *********//
	rLoader.routerHealthCheck(router, healthHandler)

}

//********* Routing API *********//

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerHealthCheck(router *gin.Engine, handler *health.V1HealthController) {
	group := router.Group("v1/check")
	group.GET("", handler.HealthCheck)
}
