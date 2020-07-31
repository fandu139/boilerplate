package errors

import (
	"github.com/sofyan48/boilerplate/src/appctx/apps"
)

type Error interface {
	New(code int, name, msg string, errs error) *apps.Error
	BadRequest() *apps.Error
	BadRequestWithErr(err error) *apps.Error
	InternalServerError(err error) *apps.Error
	DBError(err error) *apps.Error
	RecordNotFound() *apps.Error
	UnAuthorized() *apps.Error
	Forbidden(msg string) *apps.Error
}
