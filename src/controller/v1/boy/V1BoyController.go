package boy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/sofyan48/BOILERGOLANG/src/service/v1/boy"
)

// BoysController types
type V1BoysController struct {
	BoysService service.V1BoysInterface
}

// BoysController params
// @contex: gin Context
func (service *V1BoysController) BoysController(context *gin.Context) {
	result := service.BoysService.Boys()
	context.JSON(http.StatusOK, result)
	return
}
