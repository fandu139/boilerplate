package service

import (
	"github.com/sofyan48/boilerplate/src/app/v1/api/example/entity"
	"github.com/sofyan48/boilerplate/src/app/v1/api/example/repository"
)

// ExampleService ...
type ExampleService struct {
	Repository repository.ExampleRepositoryInterface
}

// ExampleServiceHandler ...
func ExampleServiceHandler() *ExampleService {
	return &ExampleService{
		Repository: repository.ExampleRepositoryHandler(),
	}
}

// ExampleServiceInterface ...
type ExampleServiceInterface interface {
	ExampleGetUUID(uuid string) (interface{}, error)
	ExampleList(pagination *entity.Pagination) (interface{}, error)
	ExampleCreateService(payload *entity.ExampleRequest) (interface{}, error)
}

// ExampleGetUUID ...
func (service *ExampleService) ExampleGetUUID(uuid string) (interface{}, error) {
	exampleData := &entity.Example{}
	err := service.Repository.GetExampleByUUID(uuid, exampleData)
	if err != nil {
		return nil, err
	}
	return exampleData, nil
}

// ExampleList ...
func (service *ExampleService) ExampleList(pagination *entity.Pagination) (interface{}, error) {
	listExample := []entity.Example{}
	listExample, err := service.Repository.GetExampleList(pagination.Limit, pagination.Page)
	if err != nil {
		return nil, err
	}
	return listExample, nil
}

// ExampleCreateService ...
func (service *ExampleService) ExampleCreateService(payload *entity.ExampleRequest) (interface{}, error) {

	return nil, nil
}
