package example

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/app/v1/api/example/controller"
	"github.com/sofyan48/boilerplate/src/middleware"
)

// RoutesExample ...
type RoutesExample struct {
	Example    controller.ExampleControllerInterface
	Middleware middleware.DefaultMiddleware
}

// RoutesExampleHandler ...
func RoutesExampleHandler() *RoutesExample {
	return &RoutesExample{
		Example: controller.ExampleControllerHandler(),
	}
}

// RoutesExampleInterface ...
type RoutesExampleInterface interface {
	Load(router *gin.Engine, version string)
}

// Load ...
func (rLoader *RoutesExample) Load(router *gin.Engine, version string) {
	group := router.Group(version + "/example")
	group.Use(rLoader.Middleware.Auth())
	group.POST("", rLoader.Example.ExampleCreate)
	group.GET("/get/:uuid", rLoader.Example.GetExampleData)
	group.GET("/list", rLoader.Example.ListExample)
	group.PUT(":uuid", rLoader.Example.UpdateExample)
	group.DELETE(":uuid", rLoader.Example.DeleteExample)
}
