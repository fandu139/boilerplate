package users

import (
	"sync"
	"time"

	"github.com/jinzhu/copier"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"

	// repositoryAPI "github.com/ralali/event-api/src/repository/api/v1"
	repository "github.com/ralali/event-api/src/repository/db/v1"
	crypto "github.com/ralali/event-api/src/util/helper/crypto"
)

// UserService derivated method
type V1UserService struct {
	userRepository repository.UserRepositoryInterface
	// userRepositoryAPI repositoryAPI.FriendAPIRepositoryInterface
}

// UserServiceHandler handler
// return: UserService
func UserServiceHandler() *V1UserService {
	return &V1UserService{
		userRepository: repository.UserRepositoryHandler(),
		// userRepositoryAPI: repositoryAPI.ThirdPartyAPIHandler(),
	}
}

// UserServiceInterface interface
type UserServiceInterface interface {
	GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse
	GetAllUser(page int, count int) []httpEntity.UserResponse
	UpdateUserByID(id int, payload httpEntity.UserRequest) bool
	InsertUsers(payload httpEntity.UserRequest) (*httpEntity.UserDetailResponse, error)
}

// GetUserByID params
// @id: int
// @waitGroup: *sync.WaitGroup
// return entity UserDetailResponse
func (service *V1UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse {
	userData := &dbEntity.Users{}
	result := &httpEntity.UserDetailResponse{}
	waitGroup.Add(1)
	go service.userRepository.GetUserByID(id, userData, waitGroup)
	waitGroup.Wait()
	copier.Copy(&result, &userData)
	return result
}

// GetAllUser params
// @id: int
// @count: int
// return entity UserResponse
func (service *V1UserService) GetAllUser(page int, count int) []httpEntity.UserResponse {
	users, _ := service.userRepository.GetUsersList(page, count)
	result := []httpEntity.UserResponse{}
	copier.Copy(&result, &users)
	return result
}

// UpdateUserByID params
// @id: int
// @payload: entity UserRequest
// return boolean
func (service *V1UserService) UpdateUserByID(id int, payload httpEntity.UserRequest) bool {
	var nows = time.Now()
	user := &dbEntity.Users{}
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	user.DateOfBirth = payload.DateOfBirth
	user.PhoneNumber = payload.PhoneNumber
	user.CurrentAddress = payload.CurrentAddress
	user.City = payload.City
	user.Province = payload.Province
	user.District = payload.District
	user.Lat = payload.Lat
	user.Lng = payload.Lng
	user.Job = payload.Job
	user.Image = payload.Image
	user.StartIncome = payload.StartIncome
	user.EndIncome = payload.EndIncome
	user.Interest = payload.Interest
	user.Habit = payload.Habit
	user.IsOrganizer = payload.IsOrganizer
	user.UpdatedAt = &nows
	err := service.userRepository.UpdateUserByID(id, user)
	if nil != err {
		return false
	}
	return true
}

// InsertUsers params
// @id: int
// @payload: entity UserRequest
// return boolean
func (service *V1UserService) InsertUsers(payload httpEntity.UserRequest) (*httpEntity.UserDetailResponse, error) {
	var nows = time.Now()
	user := &dbEntity.Users{}

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password, _ = crypto.HashPassword(payload.Password)
	user.DateOfBirth = payload.DateOfBirth
	user.PhoneNumber = payload.PhoneNumber
	user.CurrentAddress = payload.CurrentAddress
	user.City = payload.City
	user.Province = payload.Province
	user.District = payload.District
	user.Lat = payload.Lat
	user.Lng = payload.Lng
	user.Job = payload.Job
	user.Image = payload.Image
	user.StartIncome = payload.StartIncome
	user.EndIncome = payload.EndIncome
	user.Interest = payload.Interest
	user.Habit = payload.Habit
	user.IsOrganizer = payload.IsOrganizer
	user.UpdatedAt = &nows
	err := service.userRepository.InsertUsers(user)
	if err != nil {
		return nil, err
	}
	results := &httpEntity.UserDetailResponse{}
	copier.Copy(results, user)
	return results, nil
}

// ShowMyProfile params
// return UserDetailResponse
// func (service *V1UserService) ShowMyProfile() *httpEntity.UserDetailResponse {
// 	userData := &dbEntity.Users{}
// 	result := &httpEntity.UserDetailResponse{}
// 	waitGroup.Add(1)
// 	go service.userRepository.GetUserByID(id, userData, waitGroup)
// 	waitGroup.Wait()
// 	copier.Copy(&result, &userData)
// 	return result
// }
