package appctx

import (
	"net/http"

	"github.com/fandu139/boilerplate/src/appctx/apps"
	"github.com/fandu139/boilerplate/src/appctx/errors"
)

type errorContext struct{}

// AppError ...
func AppError() errors.Error {
	return &errorContext{}
}

func (e *errorContext) New(code int, name, msg string, errs error) *apps.Error {
	result := &apps.Error{}
	result.Code = code
	result.Messages = msg
	result.Name = name
	result.Errors = errs
	return result
}

func (e errorContext) BadRequest() *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusBadRequest
	result.Messages = "Check your request payload"
	result.Name = "BAD_REQUEST"
	result.Errors = nil
	return result
}

func (e errorContext) InternalServerError(err error) *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusInternalServerError
	result.Messages = "Something Went Wrong"
	result.Name = "INTERNAL_SERVICE_ERROR"
	result.Errors = err
	result.Sentry().Send()
	return result
}

func (e errorContext) Forbidden(msg string) *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusForbidden
	result.Messages = msg
	result.Name = "INTERNAL_SERVICE_FORBIDDEN"
	result.Sentry().Send()
	return result
}

func (e errorContext) DBError(err error) *apps.Error {
	switch err.Error() {
	case "record not found":
		result := &apps.Error{}
		result.Code = http.StatusNotFound
		result.Messages = "Check Your Field"
		result.Name = "DATABASE_ERROR"
		result.Errors = err
		result.Sentry().Send()
		return result
	default:
		result := &apps.Error{}
		result.Code = http.StatusInternalServerError
		result.Messages = "Something Went Wrong"
		result.Name = "DATABASE_ERROR"
		result.Errors = err
		return result
	}
}

func (e errorContext) RecordNotFound() *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusNotFound
	result.Messages = "Record Not Found"
	result.Name = "RECORD_NOT_FOUND"
	result.Errors = nil
	return result
}

func (e errorContext) UnAuthorized() *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusUnauthorized
	result.Messages = "Your Not Access"
	result.Name = "UNAUTHORIZED"
	result.Errors = nil
	return result
}

func (e errorContext) BadRequestWithErr(err error) *apps.Error {
	result := &apps.Error{}
	result.Code = http.StatusBadRequest
	result.Messages = err.Error()
	result.Name = "BAD_REQUEST"
	result.Errors = nil
	return result
}
