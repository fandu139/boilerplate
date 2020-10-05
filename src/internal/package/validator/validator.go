package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/fandu139/boilerplate/src/appctx"
	"github.com/fandu139/boilerplate/src/appctx/apps"
	"github.com/fandu139/boilerplate/src/appctx/errors"
)

type libValidator struct {
	Valid  *validator.Validate
	Errors errors.Error
}

// New ...
func New() *libValidator {
	return &libValidator{
		Valid:  validator.New(),
		Errors: appctx.AppError(),
	}
}

// Validator ...
type Validator interface {
	Request(value interface{}) *apps.Error
}

func (v *libValidator) Request(value interface{}) *apps.Error {
	// err := v.Valid.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
	// 	return len(fl.Field().String()) > 8
	// })
	// if err != nil {
	// 	return v.Errors.BadRequestWithErr(err)
	// }
	err := v.Valid.Struct(value)

	if err != nil {
		return v.Errors.BadRequestWithErr(err)
	}
	return nil
}
