package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/routes"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	V1Routers := v1.RouterLoaderHandler()
	V1Routers.Load(routers)
}
