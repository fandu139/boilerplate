package health

import (
	"github.com/jinzhu/gorm"
	"github.com/fandu139/boilerplate/src/appctx"
	"github.com/fandu139/boilerplate/src/appctx/apps"
	"github.com/fandu139/boilerplate/src/appctx/errors"
	"github.com/fandu139/boilerplate/src/appctx/result"
	"github.com/fandu139/boilerplate/src/internal/ucase/entity"
	"github.com/fandu139/boilerplate/src/internal/ucase/service"
	"github.com/fandu139/boilerplate/src/utils/cache"
	"github.com/fandu139/boilerplate/src/utils/database"
)

// Health ...
type health struct {
	DB     *gorm.DB
	Cache  cache.Contract
	Errors errors.Error
	Result result.Result
}

// New ...
func New() service.Health {
	return &health{
		DB:     database.New().MySQL().Read,
		Cache:  cache.New(),
		Errors: appctx.AppError(),
		Result: appctx.AppResult(),
	}
}

// HealthStatus ...
func (h *health) HealthStatus() (*apps.Result, *apps.Error) {
	result := &entity.HealthResponse{
		Messages: "Some Services is Failed",
	}
	err := h.DB.DB().Ping()
	if err != nil {
		return h.Result.Detail(result, "Failed Access Database"), h.Errors.DBError(err)
	}
	_, err = h.Cache.Redis().Ping()
	if err != nil {
		return h.Result.Detail(result, "Failed Access Redis"), h.Errors.DBError(err)
	}
	result.Messages = "Services OK"
	return h.Result.Detail(result, "OK"), nil
}

// HealthCheck Function
// return HealthResponse
func (h *health) HealthCheck() (*apps.Result, *apps.Error) {
	result := &entity.HealthDetailResponse{}
	result.Database.Name = "Database"
	result.Redis.Name = "Redis"

	errDB := h.DB.DB().Ping()
	if errDB != nil {
		result.Database.Health.Message = "Fail : Check Your Connection | " + errDB.Error()
	} else {
		result.Database.Health.Healthy = true
		result.Database.Health.Message = "Service OK"
	}

	_, errRD := h.Cache.Redis().Ping()
	if errRD != nil {
		result.Redis.Health.Message = "Fail : Check Your Connection | " + errRD.Error()
	} else {
		result.Redis.Health.Healthy = true
		result.Redis.Health.Message = "Service OK"
	}

	if errDB != nil {
		return h.Result.Detail(result, ""), h.Errors.DBError(errDB)
	}
	return h.Result.Detail(result, "OK"), nil
}
