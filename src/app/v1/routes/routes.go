package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/routes/docs"
	order "github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/routes/example"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/routes/health"
)

// VERSION ...
const VERSION = "v1"

// RouterLoader types
type RouterLoader struct {
	Example order.RoutesExampleInterface
	Docs    docs.RoutesDocsInterface
	Health  health.RoutesHealthInterface
}

// RouterLoaderHandler ...
func RouterLoaderHandler() *RouterLoader {
	return &RouterLoader{
		Example: order.RoutesExampleHandler(),
		Docs:    docs.RoutesDocsHandler(),
		Health:  health.RoutesHealthHandler(),
	}
}

// RouterLoaderInterface ...
type RouterLoaderInterface interface {
	Load(router *gin.Engine)
}

// Load Params
// @router: gin.Engine
func (route *RouterLoader) Load(router *gin.Engine) {
	route.Example.Load(router, VERSION)
	route.Docs.Load(router, VERSION)
	route.Health.Load(router, VERSION)

}
