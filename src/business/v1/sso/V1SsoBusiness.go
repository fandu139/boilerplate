package sso

import (
	"errors"
	"time"

	apiEntity "github.com/ralali/event-api/src/entity/api"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	repository "github.com/ralali/event-api/src/repository/db/v1"
)

//V1SsoBusiness | derivated from UserRepository
type V1SsoBusiness struct {
	UserRepository  repository.UserRepositoryInterface
	TokenRepository repository.TokenRepositoryInterface
}

// SsoBusinessHandler | Passing To sso service
func SsoBusinessHandler() *V1SsoBusiness {
	return &V1SsoBusiness{
		UserRepository:  repository.UserRepositoryHandler(),
		TokenRepository: repository.TokenRepositoryHandler(),
	}
}

// SsoBusinessInterface delcare all method
type SsoBusinessInterface interface {
	StoreGoogleUserProfile(result *dbEntity.Users) (*dbEntity.Users, error)
	StoreRalaliUserProfile(userData *apiEntity.ProfileOauthResponse) (*dbEntity.Users, error)
}

// StoreGoogleUserProfile params
// @payload: entity ResultProfileOauth
// return err
func (business *V1SsoBusiness) StoreGoogleUserProfile(result *dbEntity.Users) (*dbEntity.Users, error) {
	nows := time.Now()
	user := &dbEntity.Users{}
	user.Name = result.Name
	user.Email = result.Email
	user.Password = ""
	user.DateOfBirth = "1990-01-01"
	user.PhoneNumber = ""
	user.CurrentAddress = ""
	user.City = ""
	user.Province = ""
	user.District = ""
	user.Lat = 0
	user.Lng = 0
	user.Image = result.Image
	user.Job = ""
	user.StartIncome = 0
	user.EndIncome = 0
	user.Interest = ""
	user.Habit = ""
	user.IsOrganizer = false
	user.CreatedAt = &nows
	// isEmailAlreadyExist := business.UserRepository.CheckEmailUsers(user.Email, user)
	err := business.UserRepository.InsertUsers(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// StoreRalaliUserProfile params
// @id: int
// @payload: entity ResultProfileOauth
// return err
func (business *V1SsoBusiness) StoreRalaliUserProfile(userData *apiEntity.ProfileOauthResponse) (*dbEntity.Users, error) {
	nows := time.Now()
	user := &dbEntity.Users{}
	isEmailAlreadyExist := business.UserRepository.CheckEmailUsers(userData.Email, user)
	if isEmailAlreadyExist {
		return user, errors.New("Email Already Exist")
	}

	user.Name = userData.Name
	user.Email = userData.Email
	user.PhoneNumber = userData.Phone
	user.DateOfBirth = "1990-01-01"
	user.CurrentAddress = ""
	user.IsOrganizer = false
	user.Image = userData.ImageProfile
	user.CreatedAt = &nows
	user.UpdatedAt = &nows
	err := business.UserRepository.InsertUsers(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
