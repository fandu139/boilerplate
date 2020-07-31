package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/internal/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	rtr := routes.RouterLoaderHandler()
	rtr.Load(routers)
}
