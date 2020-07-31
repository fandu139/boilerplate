package appctx

import (
	"net/http"

	"github.com/sofyan48/boilerplate/src/appctx/apps"
	"github.com/sofyan48/boilerplate/src/appctx/result"
)

type resultContext struct{}

// AppResult ...
func AppResult() result.Result {
	return &resultContext{}
}

func (r *resultContext) New(code int, name, msg string, data interface{}) *apps.Result {
	result := &apps.Result{}
	result.Code = code
	result.Messages = msg
	result.Name = name
	result.Result = data
	return result
}

func (r *resultContext) List(data interface{}, meta interface{}) *apps.Result {
	result := &apps.Result{}
	result.Code = http.StatusOK
	result.Result = data
	result.Meta = meta
	return result
}

func (r *resultContext) Detail(data interface{}, msg string) *apps.Result {
	result := &apps.Result{}
	result.Code = http.StatusOK
	result.Result = data
	result.Messages = msg
	return result
}

func (r *resultContext) Message(name, msg string, details map[string]string) *apps.Result {
	result := &apps.Result{}
	result.Code = http.StatusOK
	result.Name = name
	result.Messages = msg
	result.MessagesDetail = details
	return result
}

func (r *resultContext) Accepted(name, msg string) *apps.Result {
	result := &apps.Result{}
	result.Code = http.StatusAccepted
	result.Name = name
	result.Messages = msg
	return result
}
