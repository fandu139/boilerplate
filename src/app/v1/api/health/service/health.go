package service

import (
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/boilerplate/src/app/v1/api/health/entity"
	"github.com/sofyan48/boilerplate/src/utils/cache"
	"github.com/sofyan48/boilerplate/src/utils/database"
)

// HealthService ...
type HealthService struct {
	DB    *gorm.DB
	Cache cache.CacheInterface
}

// HealthServiceHandler ...
func HealthServiceHandler() *HealthService {
	return &HealthService{
		DB:    database.DatabaseHandler().MySQL().Read,
		Cache: cache.CacheHandler(),
	}
}

// HealthServiceInterface ...
type HealthServiceInterface interface {
	HealthStatus() (*entity.HealthResponse, error)
	HealthCheck() (*entity.HealthDetailResponse, error)
}

// HealthStatus ...
func (service *HealthService) HealthStatus() (*entity.HealthResponse, error) {
	result := &entity.HealthResponse{
		Messages: "Some Services is Failed",
	}
	err := service.DB.DB().Ping()
	if err != nil {
		return result, err
	}
	_, err = service.Cache.Redis().Ping()
	if err != nil {
		return result, err
	}
	result.Messages = "Services OK"
	return result, nil
}

// HealthCheck Function
// return HealthResponse
func (service *HealthService) HealthCheck() (*entity.HealthDetailResponse, error) {
	result := &entity.HealthDetailResponse{}
	result.Database.Name = "Database"
	result.Redis.Name = "Redis"

	errDB := service.DB.DB().Ping()
	if errDB != nil {
		result.Database.Health.Message = "Fail : Check Your Connection | " + errDB.Error()
	} else {
		result.Database.Health.Healthy = true
		result.Database.Health.Message = "Service OK"
	}

	_, errRD := service.Cache.Redis().Ping()
	if errRD != nil {
		result.Redis.Health.Message = "Fail : Check Your Connection | " + errRD.Error()
	} else {
		result.Redis.Health.Healthy = true
		result.Redis.Health.Message = "Service OK"
	}

	if errDB != nil {
		return result, errDB
	}
	return result, nil
}
