package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/internal/routes/docs"
	"github.com/sofyan48/boilerplate/src/internal/routes/health"
)

// VERSION ...
const VERSION = "v1"

// RouterLoader types
type RouterLoader struct {
	Docs   docs.RoutesDocsInterface
	Health health.RoutesHealthInterface
}

// RouterLoaderHandler ...
func RouterLoaderHandler() *RouterLoader {
	return &RouterLoader{
		Docs:   docs.RoutesDocsHandler(),
		Health: health.RoutesHealthHandler(),
	}
}

// RouterLoaderInterface ...
type RouterLoaderInterface interface {
	Load(router *gin.Engine)
}

// Load Params
// @router: gin.Engine
func (route *RouterLoader) Load(router *gin.Engine) {
	route.Docs.Load(router, VERSION)
	route.Health.Load(router, VERSION)

}
