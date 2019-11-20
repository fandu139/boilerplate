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
	// Login
	loginHandler := &login.V1LoginController{
		LoginService: loginServices.LoginServiceHandler(),
	}

	// Sso
	ssoHandler := &sso.V1SsoController{
		SsoService: ssoServices.SsoServiceHandler(),
	}

	// health
	// Health Handler Routes
	healthHandler := &health.V1HealthController{
		HealthService: healthService.V1HealthCheckHandler(),
	}

	//********* Calling Handler To Routers *********//
	rLoader.routerLogin(router, loginHandler)
	rLoader.routerHealthCheck(router, healthHandler)
	rLoader.routerOauth(router, ssoHandler)
	rLoader.routerGoogleOauth(router, ssoHandler)

}

//********* Routing API *********//

// routerDefinition Routes for event organizer | params
// @router: gin Engine
// @handler: HealthController
func (rLoader *V1RouterLoader) routerHealthCheck(router *gin.Engine, handler *health.V1HealthController) {
	group := router.Group("v1/check")
	group.GET("", handler.HealthCheck)
}

// routerDefinition Routes for user | params
// @router: gin Engine
// @handler: UserController
func (rLoader *V1RouterLoader) routerUsers(router *gin.Engine, handler *users.V1UserController) {
	group := router.Group("v1/users")
	group.GET("", rLoader.Middleware.AuthToken(0), handler.GetUsers)
	group.GET(":id", rLoader.Middleware.AuthToken(0), handler.GetUserByID)
	group.PUT(":id", rLoader.Middleware.AuthToken(0), handler.UpdateUsersByID)
}

// routerUsersNoLogin Routes for user | params
// @router: gin Engine
// @handler: UserController
func (rLoader *V1RouterLoader) routerUsersNoLogin(router *gin.Engine, handler *users.V1UserController) {
	group := router.Group("v1/users")
	group.POST("", handler.InsertUsers)
}

// routerDefinition Routes for oauth | params
// @router: gin Engine
// @handler: SsoController
func (rLoader *V1RouterLoader) routerOauth(router *gin.Engine, handler *sso.V1SsoController) {
	group := router.Group("v1/ralali/oauth")
	group.POST("", handler.LoginRalaliOauthToken)
}

// routerDefinition Routes for oauth | params
// @router: gin Engine
// @handler: SsoController
func (rLoader *V1RouterLoader) routerGoogleOauth(router *gin.Engine, handler *sso.V1SsoController) {
	group := router.Group("v1/google/oauth")
	group.POST("", handler.GetGoogleOauthToken)
}

// routerLogin Routes for forgot | params
// @router: gin Engine
// @handler: LoginController
func (rLoader *V1RouterLoader) routerLogin(router *gin.Engine, handler *login.V1LoginController) {
	group := router.Group("v1/login")
	group.POST("", handler.Login)
	group.GET("refresh", handler.RefreshTokens)
}
