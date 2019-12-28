package users

import (
	"github.com/gin-gonic/gin"
	boy "github.com/sofyan48/BOILERGOLANG/src/controller/v1/boy"
	health "github.com/sofyan48/BOILERGOLANG/src/controller/v1/health"

	boyService "github.com/sofyan48/BOILERGOLANG/src/service/v1/boy"
	healthService "github.com/sofyan48/BOILERGOLANG/src/service/v1/health"

	"github.com/sofyan48/BOILERGOLANG/src/util/middleware"
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

	// Health Handler Routes
	boyHandler := &boy.V1BoysController{
		BoysService: boyService.V1BoysHandler(),
	}

	//********* Calling Handler To Routers *********//
	rLoader.routerHealthCheck(router, healthHandler)
	rLoader.routerBoys(router, boyHandler)

}

//********* Routing API *********//

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerHealthCheck(router *gin.Engine, handler *health.V1HealthController) {
	group := router.Group("v1/check")
	group.GET("", handler.HealthCheck)
}

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerBoys(router *gin.Engine, handler *boy.V1BoysController) {
	group := router.Group("v1/boy")
	group.GET("", handler.BoysController)
}
