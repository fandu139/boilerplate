package v1

import (
	"sync"

	"github.com/jinzhu/gorm"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	connection "github.com/ralali/event-api/src/util/helper/mysqlconnection"
)

// LoginRepository type
type LoginRepository struct {
	DB gorm.DB
}

// LoginRepositoryHandler Users handler repo
// return: LoginRepository
func LoginRepositoryHandler() *LoginRepository {
	return &LoginRepository{DB: *connection.GetConnection()}
}

// LoginRepositoryInterface interface
type LoginRepositoryInterface interface {
	GetUserByEmail(email string, userData *dbEntity.Users, wg *sync.WaitGroup) error
}

// GetUserByEmail params
// @id: int
// @userData: entity Users
// wg *sync.WaitGroup
// return error
func (repository *LoginRepository) GetUserByEmail(email string, userData *dbEntity.Users, wg *sync.WaitGroup) error {
	query := repository.DB.Table("users")
	query = query.Where("email=?", email)
	query = query.First(userData)
	wg.Done()
	return query.Error
}
