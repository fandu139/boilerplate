package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/boilerplate/src/internal/package/rest"
	"github.com/sofyan48/boilerplate/src/internal/ucase/contract"
	"github.com/sofyan48/boilerplate/src/internal/ucase/service"
	"github.com/sofyan48/boilerplate/src/internal/ucase/service/health"
)

// Health types
type handler struct {
	Service service.Health
	Rest    rest.RestInterface
}

// New ...
func New() contract.Health {
	return &handler{
		Service: health.New(),
		Rest:    rest.RestHandler(),
	}
}

// HealthStatus example
// @Summary Check Health Status
// @Description get string by ID
// @Tags Health Status
// @Accept  json
// @Produce  json
// @Success 200 {array} rest.Result
// @Failure 400 {object} rest.ErrorResponse
// @Failure 404 {object} rest.ErrorResponse
// @Failure 500 {object} rest.ErrorResponse
// @Router /v1/health [get]
func (h *handler) HealthStatus(context *gin.Context) {
	result, err := h.Service.HealthStatus()
	if err != nil {
		h.Rest.ErrorResponse(err)
		return
	}
	h.Rest.Result(result)
	return
}

// HealthCheck detail response check all service dependecies
// @Summary Check Health Status Detail
// @Description Check Health Status Detail
// @Tags HealthCheck Status
// @Accept  json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Success 200 {array} rest.Result
// @Failure 400 {object} rest.ErrorResponse
// @Failure 404 {object} rest.ErrorResponse
// @Failure 500 {object} rest.ErrorResponse
// @Router /v1/health/check [get]
func (h *handler) HealthCheck(context *gin.Context) {
	result, err := h.Service.HealthCheck()
	if err != nil {
		h.Rest.ErrorResponse(err)
		return
	}
	h.Rest.Result(result)
	return
}
