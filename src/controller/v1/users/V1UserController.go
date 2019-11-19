package users

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	services "github.com/ralali/event-api/src/service/v1/users"
	"github.com/ralali/event-api/src/util/helper/rest"
)

// Limitofset types
type Limitofset struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// V1UserController types
type V1UserController struct {
	UserService services.UserServiceInterface
}

// GetUserByID params
// @context: gin context
func (service *V1UserController) GetUserByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
	}
	result := service.UserService.GetUserByID(id, &sync.WaitGroup{})
	if result == nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Data Null")
	}
	rest.ResponseData(context, http.StatusOK, result)
}

// GetUsers params
// @context: gin context
func (service *V1UserController) GetUsers(context *gin.Context) {
	queryparam := Limitofset{}
	err := context.ShouldBindQuery(&queryparam)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	result := service.UserService.GetAllUser(queryparam.Limit, queryparam.Offset)
	rest.ResponseList(context, http.StatusOK, result)
	return
}

// UpdateUsersByID params
// @context: gin context
func (service *V1UserController) UpdateUsersByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	service.UserService.UpdateUserByID(id, payload)
	rest.ResponseMessages(context, http.StatusNoContent, "Edited")
	return
}

// InsertUsers params
// @context: gin context
func (service *V1UserController) InsertUsers(context *gin.Context) {
	payload := httpEntity.UserRequest{}
	if err := context.ShouldBind(&payload); err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, "Bad Request")
		return
	}
	fmt.Println(payload)
	results, err := service.UserService.InsertUsers(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusUnauthorized, fmt.Sprintf("%e", err))
		return
	}
	rest.ResponseData(context, http.StatusCreated, results)
	return
}
