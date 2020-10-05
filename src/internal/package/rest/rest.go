package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/fandu139/boilerplate/src/appctx/apps"
)

// APIResponse types
type APIResponse struct {
	Code          int               `json:"code"`
	Name          string            `json:"name,omitempty"`
	Error         string            `json:"error,omitempty"`
	Meta          interface{}       `json:"meta,omitempty"`
	Results       interface{}       `json:"data,omitempty"`
	Message       string            `json:"message"`
	MessageDetail map[string]string `json:"message_detail,omitempty"`
}

var context *gin.Context

type Rest struct {
	Context *gin.Context
}

func RestHandler() *Rest {
	return &Rest{}
}

type RestInterface interface {
	// Result(context *gin.Context, result *apps.Result)
	Result(result *apps.Result)
	ErrorResponse(errs *apps.Error)
}

// Setup ...
func Setup(ctx *gin.Context) {
	context = ctx
}

// Result ...
// func (r *Rest) Result(context *gin.Context, result *apps.Result) {
func (r *Rest) Result(result *apps.Result) {
	response := APIResponse{
		Code:          result.Code,
		Meta:          result.Meta,
		Results:       result.Result,
		Message:       result.Messages,
		MessageDetail: result.MessagesDetail,
	}
	context.JSON(result.Code, response)
}

// ErrorResponse ...
func (r *Rest) ErrorResponse(errs *apps.Error) {
	response := APIResponse{
		Code:    errs.Code,
		Name:    errs.Name,
		Error:   errs.Errors.Error(),
		Message: errs.Messages,
	}
	context.JSON(errs.Code, response)
	// return
}
