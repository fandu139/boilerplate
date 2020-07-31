package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sofyan48/boilerplate/src/internal/ucase/entity"
)

// Example interface
type Example interface {
	GetByID(id int, exampleData *entity.Example) error
	GetList(limit int, offset int) ([]entity.Example, error)
	GetByUUID(uuid string, exampleStatusData *entity.Example) error
	InsertExample(tx *gorm.DB, usersData *entity.Example) error
}
