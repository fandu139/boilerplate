package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/boilerplate/src/app/v1/api/example/entity"
	"github.com/sofyan48/boilerplate/src/utils/database"
)

// TABLE ...
const TABLE = "tb_example"

// ExampleRepository types
type ExampleRepository struct {
	DB *gorm.DB
}

// ExampleRepositoryHandler Example handler repo
// return: ExampleRepository
func ExampleRepositoryHandler() *ExampleRepository {
	return &ExampleRepository{
		DB: database.DatabaseHandler().MySQL().Read,
	}
}

// ExampleRepositoryInterface interface
type ExampleRepositoryInterface interface {
	GetExampleByID(id int, exampleData *entity.Example) error
	GetExampleList(limit int, offset int) ([]entity.Example, error)
	GetExampleByUUID(uuid string, exampleStatusData *entity.Example) error
	InsertExample(TX *gorm.DB, usersData *entity.Example) error
}

// GetExampleByUUID params
// @id: int
// @exampleData: entity Example
// wg *sync.WaitGroup
// return error
func (repository *ExampleRepository) GetExampleByUUID(uuid string, exampleStatusData *entity.Example) error {
	query := repository.DB.Table(TABLE)
	query = query.Where("uuid=?", uuid)
	query = query.First(&exampleStatusData)
	return query.Error
}

// GetExampleByID params
// @id: int
// @exampleData: entity Example
// wg *sync.WaitGroup
// return error
func (repository *ExampleRepository) GetExampleByID(id int, exampleData *entity.Example) error {
	query := repository.DB.Table(TABLE)
	query = query.Where("id_example=? AND deleted_at IS NULL", id)
	query = query.First(&exampleData)
	return query.Error
}

// GetExampleList params
// @id: int
// @exampleData: entity Example
// return entity,error
func (repository *ExampleRepository) GetExampleList(limit int, offset int) ([]entity.Example, error) {
	users := []entity.Example{}
	query := repository.DB.Table(TABLE)
	query = query.Where("deleted_at IS NULL")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	return users, query.Error
}

// InsertExample params
// @exampleData: entity Example
// return error
func (repository *ExampleRepository) InsertExample(TX *gorm.DB, usersData *entity.Example) error {
	query := TX.Table(TABLE)
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}
