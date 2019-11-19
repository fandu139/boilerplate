package sso

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	apiEntity "github.com/ralali/event-api/src/entity/api"
	service "github.com/ralali/event-api/src/service/v1/sso"
	"github.com/ralali/event-api/src/util/helper/rest"
	"github.com/ralali/event-api/src/util/middleware"
)

type V1SsoController struct {
	SsoService service.SsoServiceInterface
	Middleware middleware.DefaultMiddleware
}

// GetGoogleOauthToken params
// @contex: gin Context
func (service *V1SsoController) GetGoogleOauthToken(context *gin.Context) {
	payload := apiEntity.SsoGoogleRequest{}
	err := context.ShouldBind(&payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	// fmt.Println(payload)
	results, err := service.SsoService.GetGoogleOauthToken(&payload, &sync.WaitGroup{})

	if err != nil {
		rest.ResponseMessages(context, http.StatusUnauthorized, err.Error())
		return
	}

	rest.ResponseData(context, http.StatusOK, results)
	return

}

// LoginRalaliOauthToken params
// @contex: gin Context
func (service *V1SsoController) LoginRalaliOauthToken(context *gin.Context) {
	payload := apiEntity.SsoAuthByPassword{}
	if err := context.ShouldBindJSON(&payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	stats, err := service.SsoService.AuthRalaliLogin(&payload, &sync.WaitGroup{})
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}
	context.JSON(http.StatusOK, stats)
	return
}
