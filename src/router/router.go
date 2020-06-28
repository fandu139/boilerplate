package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sofyan48/boilerplate/src/app/v1/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	V1Routers := v1.RouterLoaderHandler()
	V1Routers.Load(routers)
}
