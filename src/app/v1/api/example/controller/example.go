package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/api/example/entity"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/api/example/service"
	"github.com/orn-id/orn-mn-boilerplate-go/src/app/v1/package/rest"
)

// ExampleController ...
type ExampleController struct {
	Service service.ExampleServiceInterface
}

// ExampleControllerHandler ...
func ExampleControllerHandler() *ExampleController {
	return &ExampleController{
		Service: service.ExampleServiceHandler(),
	}
}

// ExampleControllerInterface ...
type ExampleControllerInterface interface {
	ExampleCreate(context *gin.Context)
	UpdateExample(context *gin.Context)
	GetExampleData(context *gin.Context)
	DeleteExample(context *gin.Context)
	ListExample(context *gin.Context)
}

// ExampleCreate ...
func (ctrl *ExampleController) ExampleCreate(context *gin.Context) {
	payload := &entity.ExampleRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.ExampleCreateService(payload)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, nil, "OK")
	return
}

// GetExampleData ...
func (ctrl *ExampleController) GetExampleData(context *gin.Context) {
	uuid := context.Param("uuid")
	result, err := ctrl.Service.ExampleGetUUID(uuid)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, nil, "Success")
	return
}

// ListExample ...
func (ctrl *ExampleController) ListExample(context *gin.Context) {
	pagination := &entity.Pagination{}
	context.ShouldBind(pagination)
	result, err := ctrl.Service.ExampleList(pagination)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, pagination, "Success")
	return
}

// UpdateExample ...
func (ctrl *ExampleController) UpdateExample(context *gin.Context) {
	rest.SuccessResponse(context, "OK", nil, nil)
}

// DeleteExample ...
func (ctrl *ExampleController) DeleteExample(context *gin.Context) {
	rest.SuccessResponse(context, "OK", nil, nil)
}
