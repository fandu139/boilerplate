package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/api/health/service"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/package/rest"
)

// HealthController types
type HealthController struct {
	Service service.HealthServiceInterface
}

// HealthControllerHandler ...
func HealthControllerHandler() *HealthController {
	return &HealthController{
		Service: service.HealthServiceHandler(),
	}
}

// HealthControllerInterface ...
type HealthControllerInterface interface {
	HealthStatus(context *gin.Context)
	HealthCheck(context *gin.Context)
}

// HealthStatus example
// @Summary Check Health Status
// @Description get string by ID
// @Tags Health Status
// @Accept  json
// @Produce  json
// @Success 200 {array} rest.RestResponseData
// @Failure 400 {object} rest.RestResponseData
// @Failure 404 {object} rest.RestResponseData
// @Failure 500 {object} rest.RestResponseData
// @Router /v1/health [get]
func (handler *HealthController) HealthStatus(context *gin.Context) {
	result, err := handler.Service.HealthStatus()
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, nil, "Service OK")
	return
}

// HealthCheck detail response check all service dependecies
// @Summary Check Health Status Detail
// @Description Check Health Status Detail
// @Tags HealthCheck Status
// @Accept  json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Success 200 {array} rest.RestResponseData
// @Failure 400 {object} rest.RestResponseData
// @Failure 404 {object} rest.RestResponseData
// @Failure 500 {object} rest.RestResponseData
// @Router /v1/health/check [get]
func (handler *HealthController) HealthCheck(context *gin.Context) {
	result, err := handler.Service.HealthCheck()

	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, nil, "Service OK")
	return
}
