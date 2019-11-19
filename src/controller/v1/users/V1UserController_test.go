package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"

	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	"github.com/stretchr/testify/assert"
)

type UserServiceMock struct{}

func (service *UserServiceMock) GetUserByID(id int, wg *sync.WaitGroup) *httpEntity.UserDetailResponse {
	return &httpEntity.UserDetailResponse{
		ID:   uint(id),
		Name: "Test Name Mock",
	}
}

func (service *UserServiceMock) GetAllUser(page int, count int) []httpEntity.UserResponse {
	users := []httpEntity.UserResponse{}
	const (
		ID1 = iota + 1
		ID2
		ID3
	)
	users = append(users, httpEntity.UserResponse{
		ID:   uint(ID1),
		Name: "Test Name Mock",
	})
	users = append(users, httpEntity.UserResponse{
		ID:   uint(ID2),
		Name: "Test Name Mock",
	})
	users = append(users, httpEntity.UserResponse{
		ID:   uint(ID3),
		Name: "Test Name Mock",
	})
	return users
}

func (service *UserServiceMock) UpdateUserByID(id int, payload httpEntity.UserRequest) bool {
	return true
}

func TestUserGetByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 1
	url := "/v1/users/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := httpEntity.UserDetailResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(res.ID, idTest, "It should be same ID")
	assert.Equal(res.Name, "Test Name Mock", "It should be same Name")
}

func TestGetUserListMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	limit := 1
	offset := 100

	url := "/v1/users?limit=" + fmt.Sprint(limit) + "&offset=" + fmt.Sprint(offset)
	c.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	r.ServeHTTP(resp, c.Request)
	assert.Equal(http.StatusOK, resp.Code, "Status should be 200")

	res := []httpEntity.UserResponse{}
	err := json.Unmarshal([]byte(resp.Body.String()), &res)

	assert.Equal(err, nil, "should have no error")
	assert.Equal(len(res) >= 0, true, "length must in minimum value")
	assert.Equal(len(res) == 3, true, "length value must match with mock data")
}

func TestUpdateUserByIDMock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	c, r, resp := LoadRouterTestMock()

	var idTest uint = 1
	payload := &httpEntity.UserRequest{
		Name: "Test Name",
	}
	jsonPayload, _ := json.Marshal(payload)
	url := "/v1/users/" + fmt.Sprint(idTest)
	c.Request, _ = http.NewRequest(http.MethodPut, url, bytes.NewBufferString(string(jsonPayload)))
	c.Request.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(resp, c.Request)

	assert.Equal(http.StatusNoContent, resp.Code, "Status should be 204")
}
