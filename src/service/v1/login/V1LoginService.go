package login

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	repository "github.com/ralali/event-api/src/repository/db/v1"
	crypto "github.com/ralali/event-api/src/util/helper/crypto"
	cache "github.com/ralali/event-api/src/util/helper/redis"
	"github.com/ralali/event-api/src/util/middleware"
)

// LoginService derivated method
type LoginService struct {
	loginRepository   repository.LoginRepositoryInterface
	Cache             redis.Conn
	DefaultMiddleWare middleware.DefaultMiddleware
}

// LoginServiceHandler handler
// return: LoginService
func LoginServiceHandler() *LoginService {
	return &LoginService{
		loginRepository: repository.LoginRepositoryHandler(),
		Cache:           cache.GetConnection(),
	}
}

// LoginServiceInterface interface
type LoginServiceInterface interface {
	LoginControll(payload httpEntity.LoginRequest, waitGroup *sync.WaitGroup) (*httpEntity.LoginDetailResponse, error)
	RefreshToken(context *gin.Context) (*httpEntity.LoginDetailResponse, error)
}

// LoginControll params
// @email: string
// @waitGroup: *sync.WaitGroup
// return entity UserDetailResponse
func (service *LoginService) LoginControll(payload httpEntity.LoginRequest, waitGroup *sync.WaitGroup) (*httpEntity.LoginDetailResponse, error) {
	email := payload.Email
	user := &dbEntity.Users{}
	res := &httpEntity.LoginDetailResponse{}

	waitGroup.Add(1)
	go service.loginRepository.GetUserByEmail(email, user, waitGroup)
	waitGroup.Wait()

	passCheck := crypto.CheckPasswordHash(payload.Password, user.Password)
	if !passCheck {
		err := errors.New("Email Or Password Not Valid")
		return res, err
	}
	sessionDataUsers := httpEntity.SessionDataUsers{}
	sessionDataUsers.ID = user.ID
	sessionDataUsers.Name = user.Name
	sessionDataUsers.Email = user.Email
	sessionDataUsers.Image = user.Image
	sessionDataUsers.IsOrganizers = user.IsOrganizer
	expirationTime := time.Unix(1605194980, 0)

	sessdata, _ := json.Marshal(sessionDataUsers)

	claims := &httpEntity.Claims{}
	claims.SessionData = string(sessdata)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    user.Email,
		Id:        string(user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return res, err
	}

	res.IsRegistered = true
	res.Token = tokenString
	res.UserData = sessionDataUsers
	res.Expired = expirationTime.Unix()

	return res, nil
}

// LogoutControll params
// @token: string
func (service *LoginService) LogoutControll(token string) {
	redis.String(cache.Store.Do("DEL", token))
}

// RefreshToken params
func (service *LoginService) RefreshToken(context *gin.Context) (*httpEntity.LoginDetailResponse, error) {
	res := &httpEntity.LoginDetailResponse{}
	token := service.DefaultMiddleWare.GetToken(context)
	claims := &httpEntity.Claims{}
	newToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if !newToken.Valid {
		return res, err
	}
	if err != nil {
		return res, err
	}
	expirationTime := time.Unix(1605194980, 0)
	claims.ExpiresAt = expirationTime.Unix()
	tokenNew := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenNew.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return res, err
	}
	res.IsRegistered = true
	res.Token = tokenString
	res.Expired = expirationTime.Unix()
	return res, nil
}
