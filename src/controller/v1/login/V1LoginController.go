package login

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	services "github.com/ralali/event-api/src/service/v1/login"
	"github.com/ralali/event-api/src/util/helper/rest"
)

// V1LoginController types
type V1LoginController struct {
	LoginService services.LoginServiceInterface
}

// Login params
// @context: gin context
func (ctrl *V1LoginController) Login(context *gin.Context) {
	payload := httpEntity.LoginRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		errorstr := fmt.Sprintf("%e", err)
		rest.ResponseMessages(context, http.StatusUnauthorized, "Bad Request | "+errorstr)
		return
	}
	result, err := ctrl.LoginService.LoginControll(payload, &sync.WaitGroup{})
	if err != nil {
		rest.ResponseMessages(context, http.StatusUnauthorized, fmt.Sprintf("%e", err))
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// RefreshTokens params
// @context: *gin.Context
func (ctrl *V1LoginController) RefreshTokens(context *gin.Context) {
	res, err := ctrl.LoginService.RefreshToken(context)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, fmt.Sprintf("%e", err))
		return
	}
	rest.ResponseData(context, http.StatusOK, res)
	return
}
