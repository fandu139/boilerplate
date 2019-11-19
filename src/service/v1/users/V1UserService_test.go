package users

import (
	"fmt"
	"sync"
	"testing"
	"time"

	modelDB "github.com/ralali/event-api/src/entity/db/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// SETUP
type repositoryDBMock struct {
	mock.Mock
}

func (repository *repositoryDBMock) GetUserByID(id int, userData *modelDB.Users, wg *sync.WaitGroup) error {
	var nows = time.Now()
	repository.Called(id, userData, wg)
	userData.ID = uint(id)
	userData.Name = "Test Name"
	userData.Email = "agus@ralali.com"
	userData.Password = "1234"
	userData.DateOfBirth = "2019-01-01"
	userData.PhoneNumber = "0128121"
	userData.CurrentAddress = ""
	userData.City = ""
	userData.Province = ""
	userData.District = ""
	userData.Lat = 1.000000
	userData.Lng = 2.000000
	userData.Job = ""
	userData.Image = "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg"
	userData.StartIncome = 0
	userData.EndIncome = 0
	userData.Interest = "OK"
	userData.Habit = "OK"
	userData.IsOrganizer = true
	userData.CreatedAt = &nows
	userData.UpdatedAt = &nows
	userData.DeletedAt = &nows
	wg.Done()
	return nil
}

func (repository *repositoryDBMock) UpdateUserByID(id int, userData *modelDB.Users) error {
	var nows = time.Now()
	repository.Called(id, userData)
	userData.ID = uint(id)
	userData.Name = fmt.Sprintf("Updated - %s", "Test Name")
	userData.Email = "agus@ralali.com"
	userData.Password = "1234"
	userData.DateOfBirth = "2019-01-01"
	userData.PhoneNumber = "0128121"
	userData.CurrentAddress = ""
	userData.City = ""
	userData.Province = ""
	userData.District = ""
	userData.Lat = 1.000000
	userData.Lng = 2.000000
	userData.Job = ""
	userData.Image = "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg"
	userData.StartIncome = 0
	userData.EndIncome = 0
	userData.Interest = "OK"
	userData.Habit = "OK"
	userData.IsOrganizer = true
	userData.CreatedAt = &nows
	userData.UpdatedAt = &nows
	userData.DeletedAt = &nows
	return nil
}

func (repository *repositoryDBMock) GetUsersList(limit int, offset int) ([]modelDB.Users, error) {
	var nows = time.Now()
	repository.Called(limit, offset)
	users := []modelDB.Users{}
	const (
		ID1 = iota + 1
		ID2 = iota + 2
		ID3 = iota + 3
	)
	users = append(users, modelDB.Users{
		ID:             uint(ID1),
		Name:           "Test Name",
		Email:          "agus@ralali.com",
		Password:       "1234",
		DateOfBirth:    "2019-01-01",
		PhoneNumber:    "0128121",
		CurrentAddress: "",
		City:           "",
		Province:       "",
		District:       "",
		Lat:            1.000000,
		Lng:            2.000000,
		Job:            "",
		Image:          "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg",
		StartIncome:    0,
		EndIncome:      0,
		Interest:       "OK",
		Habit:          "OK",
		IsOrganizer:    true,
		CreatedAt:      &nows,
		UpdatedAt:      &nows,
		DeletedAt:      &nows,
	})
	users = append(users, modelDB.Users{
		ID:             uint(ID2),
		Name:           "Test Name",
		Email:          "agus@ralali.com",
		Password:       "1234",
		DateOfBirth:    "2019-01-01",
		PhoneNumber:    "0128121",
		CurrentAddress: "",
		City:           "",
		Province:       "",
		District:       "",
		Lat:            1.000000,
		Lng:            2.000000,
		Job:            "",
		Image:          "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg",
		StartIncome:    0,
		EndIncome:      0,
		Interest:       "OK",
		Habit:          "OK",
		IsOrganizer:    true,
		CreatedAt:      &nows,
		UpdatedAt:      &nows,
		DeletedAt:      &nows,
	})
	users = append(users, modelDB.Users{
		ID:             uint(ID3),
		Name:           "Test Name",
		Email:          "agus@ralali.com",
		Password:       "1234",
		DateOfBirth:    "2019-01-01",
		PhoneNumber:    "0128121",
		CurrentAddress: "",
		City:           "",
		Province:       "",
		District:       "",
		Lat:            1.000000,
		Lng:            2.000000,
		Job:            "",
		Image:          "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg",
		StartIncome:    0,
		EndIncome:      0,
		Interest:       "OK",
		Habit:          "OK",
		IsOrganizer:    true,
		CreatedAt:      &nows,
		UpdatedAt:      &nows,
		DeletedAt:      &nows,
	})
	return users, nil
}

func (repository *repositoryDBMock) CheckEmailUsers(email string, userData *modelDB.Users) bool {
	var nows = time.Now()
	repository.Called(email, userData)
	userData.ID = 1
	userData.Name = "Test Name"
	userData.Email = email
	userData.Password = "1234"
	userData.DateOfBirth = "2019-01-01"
	userData.PhoneNumber = "0128121"
	userData.CurrentAddress = ""
	userData.City = ""
	userData.Province = ""
	userData.District = ""
	userData.Lat = 1.000000
	userData.Lng = 2.000000
	userData.Job = ""
	userData.Image = "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg"
	userData.StartIncome = 0
	userData.EndIncome = 0
	userData.Interest = "OK"
	userData.Habit = "OK"
	userData.IsOrganizer = true
	userData.CreatedAt = &nows
	userData.UpdatedAt = &nows
	userData.DeletedAt = &nows
	return true
}

func (repository *repositoryDBMock) InsertUsers(userData *modelDB.Users) error {
	var nows = time.Now()
	repository.Called(userData)
	userData.ID = 1
	userData.Name = "Test Name"
	userData.Email = "email"
	userData.Password = "1234"
	userData.DateOfBirth = "2019-01-01"
	userData.PhoneNumber = "0128121"
	userData.CurrentAddress = ""
	userData.City = ""
	userData.Province = ""
	userData.District = ""
	userData.Lat = 1.000000
	userData.Lng = 2.000000
	userData.Job = ""
	userData.Image = "image-9b395d74-a233-4d95-b539-dd3057d8fd43jpg-5b725f9f287a1997341417.jpg"
	userData.StartIncome = 0
	userData.EndIncome = 0
	userData.Interest = "OK"
	userData.Habit = "OK"
	userData.IsOrganizer = true
	userData.CreatedAt = &nows
	userData.UpdatedAt = &nows
	userData.DeletedAt = &nows
	return nil
}

type repositoryAPIMock struct {
	mock.Mock
}

// TEST
func TestUserServiceGetUserByIDMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryDBMock{}
	wg := &sync.WaitGroup{}

	user := &modelDB.Users{}
	testID := 1
	dbMockData.On("GetUserByID", testID, user, wg).Return(nil)

	V1UserService := V1UserService{&dbMockData}
	resultFuncService := V1UserService.GetUserByID(testID, wg)
	assert.Equal(t, uint(testID), resultFuncService.ID, "It should be same ID")
}

func TestUserServiceGetAllUserMocked(t *testing.T) {
	t.Parallel()
	dbMockData := repositoryDBMock{}
	limit := 1
	offset := 3
	dbMockData.On("GetUsersList", limit, offset).Return([]modelDB.Users{}, nil)
	V1UserService := V1UserService{&dbMockData}
	resultFuncService := V1UserService.GetAllUser(limit, offset)
	assert.Equal(t, len(resultFuncService), 3, "It should be same length as Mock Data")
}

func TestUserServiceUpdateUserByIDMocked(t *testing.T) {
	// t.Parallel()
	// dbMockData := repositoryDBMock{}
	// var testId int = 1
	// dbMockData.On("UpdateUserByID", testId, &modelDB.Users{
	// 	ID:          1,
	// 	Name:        "Test Name",
	// 	Email:       "agus@ralali.com",
	// 	Password:    "1234",
	// 	DateOfBirth: "2019-01-01",
	// 	PhoneNumber: "0128121",
	// }).Return(nil)
	// userService := V1UserService{&dbMockData}
	// resultFuncService := userService.UpdateUserByID(testId, modelHttp.UserRequest{
	// 	Name:        "Test Update",
	// 	Email:       "agus@ralali.com",
	// 	Password:    "1234",
	// 	DateOfBirth: "2019-01-01",
	// 	PhoneNumber: "0128121",
	// })
	// fmt.Println(resultFuncService)
	// assert.Equal(t, resultFuncService, true, "It should be true")
	// assert.Equal(t, resultFuncService, true, "It should be true")
}
