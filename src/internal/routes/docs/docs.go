package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/fandu139/boilerplate/src/internal/swagger"
)

// RoutesDocs ...
type RoutesDocs struct {
}

// RoutesOrderHandler ..
func RoutesDocsHandler() *RoutesDocs {
	return &RoutesDocs{}
}

// RoutesOrderInterface ...
type RoutesDocsInterface interface {
	Load(router *gin.Engine, version string)
}

// Load ..
func (rLoader *RoutesDocs) Load(router *gin.Engine, version string) {
	group := router.Group(version + "/docs")
	url := ginSwagger.URL("swagger/doc.json")
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
