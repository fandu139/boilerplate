package example

import (
	"github.com/jinzhu/gorm"

	"github.com/sofyan48/boilerplate/src/internal/ucase/entity"
	"github.com/sofyan48/boilerplate/src/internal/ucase/repository"
	"github.com/sofyan48/boilerplate/src/utils/database"
)

// TABLE ...
const TABLE = "tb_example"

// example types
type example struct {
	Read *gorm.DB
}

// New ...
func New() repository.Example {
	return &example{
		Read: database.New().MySQL().Read,
	}
}

// GetByUUID params
// @id: int
// @exampleData: entity Example
// wg *sync.WaitGroup
// return error
func (repository *example) GetByUUID(uuid string, exampleStatusData *entity.Example) error {
	query := repository.Read.Table(TABLE)
	query = query.Where("uuid=?", uuid)
	query = query.First(&exampleStatusData)
	query.Close()
	return query.Error
}

// GetByID params
// @id: int
// @exampleData: entity Example
// wg *sync.WaitGroup
// return error
func (repository *example) GetByID(id int, exampleData *entity.Example) error {
	query := repository.Read.Table(TABLE)
	query = query.Where("id_example=? AND deleted_at IS NULL", id)
	query = query.First(&exampleData)
	query.Close()
	return query.Error
}

// GetList params
// @id: int
// @exampleData: entity Example
// return entity,error
func (repository *example) GetList(limit int, offset int) ([]entity.Example, error) {
	users := []entity.Example{}
	query := repository.Read.Table(TABLE)
	query = query.Where("deleted_at IS NULL")
	query = query.Limit(limit).Offset(offset)
	query = query.Find(&users)
	query.Close()
	return users, query.Error
}

// InsertExample params
// @exampleData: entity Example
// return error
func (repository *example) InsertExample(tx *gorm.DB, usersData *entity.Example) error {
	query := tx.Table(TABLE)
	query = query.Create(usersData)
	query.Scan(&usersData)
	return query.Error
}
