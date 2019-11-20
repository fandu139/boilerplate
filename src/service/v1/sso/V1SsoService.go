package sso

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	// "reflect"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/copier"
	business "github.com/ralali/event-api/src/business/v1/sso"
	apiEntity "github.com/ralali/event-api/src/entity/api"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	repository "github.com/ralali/event-api/src/repository/db/v1"
	cache "github.com/ralali/event-api/src/util/helper/redis"
	"github.com/ralali/event-api/src/util/helper/str_process"
	"github.com/ralali/event-api/src/util/middleware"

	"google.golang.org/api/oauth2/v2"
)

// V1SsoService | Derivated from UserRepository
type V1SsoService struct {
	UserRepository    repository.UserRepositoryInterface
	SsoBusiness       business.SsoBusinessInterface
	Cache             redis.Conn
	DefaultMiddleware middleware.DefaultMiddleware
}

// SsoServiceHandler func
// return SsoService
func SsoServiceHandler() *V1SsoService {
	return &V1SsoService{
		UserRepository: repository.UserRepositoryHandler(),
		SsoBusiness:    business.SsoBusinessHandler(),
		Cache:          cache.GetConnection(),
	}
}

//SsoServiceInterface declare All Method
type SsoServiceInterface interface {
	GetGoogleOauthToken(payload *apiEntity.SsoGoogleRequest, waitGroup *sync.WaitGroup) (interface{}, error)
	GetRalaliOauthToken() (*apiEntity.SsoResponseBeforeUsers, []byte)
	CreateClientAuthRalali() (*apiEntity.SsoRalaliClientCreate, []byte)
	AuthRalaliLogin(payload *apiEntity.SsoAuthByPassword, waitGroup *sync.WaitGroup) (interface{}, error)
}

// GetGoogleOauthToken params
// @payload: SsoGoogleRequest
// return SsoGoogleResponse
func (service *V1SsoService) GetGoogleOauthToken(payload *apiEntity.SsoGoogleRequest, waitGroup *sync.WaitGroup) (interface{}, error) {
	googleData := &apiEntity.GoogleReturnResponse{}
	result := &httpEntity.LoginOauthDetailResponse{}
	detailResUser := httpEntity.UserDetailResponse{}

	// Verify IDToken from Google
	var httpClient = &http.Client{}
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(payload.IDToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	copier.Copy(&googleData, &tokenInfo)
	if !googleData.VerifiedEmail {
		return nil, err
	}

	user := &dbEntity.Users{}
	user.Name = payload.Name
	user.Email = payload.Email
	user.Image = payload.Image
	isEmailAlreadyExist := service.UserRepository.CheckEmailUsers(user.Email, user)
	usersData := &dbEntity.Users{}
	if !isEmailAlreadyExist {
		usersData, err = service.SsoBusiness.StoreGoogleUserProfile(user)
		if err != nil {
			return nil, err
		}
	}

	sessionDataUsers := httpEntity.SessionDataUsers{}

	sessionDataUsers.ID = usersData.ID
	sessionDataUsers.Name = usersData.Name
	sessionDataUsers.Email = usersData.Email
	sessionDataUsers.IsOrganizers = usersData.IsOrganizer
	expirationTime := time.Unix(1605194980, 0)

	sessdata, _ := json.Marshal(sessionDataUsers)
	claims := &httpEntity.Claims{}
	claims.SessionData = string(sessdata)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    usersData.Email,
		Id:        string(usersData.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	result.Token = tokenString
	result.SessionData = sessionDataUsers

	copier.Copy(&detailResUser, &usersData)
	result.IsRegistered = isEmailAlreadyExist
	result.UserData = detailResUser
	result.SessionData.Image = payload.Image
	result.Expired = expirationTime.Unix()
	return result, nil
}

// CreateClientAuthRalali params
// return SsoResponse
func (service *V1SsoService) CreateClientAuthRalali() (*apiEntity.SsoRalaliClientCreate, []byte) {
	data := &apiEntity.ClientCreateRalaliSso{}
	result := &apiEntity.SsoRalaliClientCreate{}

	data.ClientName = "BIG-" + str_process.GenerateUniqueCode("EVENT")
	data.IsFirtsParty = true
	data.RedirectUris = []string{os.Getenv("RALALI_REDIRECT_URI")}
	jsonStr, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", os.Getenv("RALALI_OAUTH_URL")+"/v2/client", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+os.Getenv("RALALI_AUTHORIZATION"))
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return result, body
	}
	json.Unmarshal([]byte(body), &result)
	return result, nil
}

// GetRalaliOauthToken params
// return SsoResponse
func (service *V1SsoService) GetRalaliOauthToken() (*apiEntity.SsoResponseBeforeUsers, []byte) {
	result := &apiEntity.SsoResponseBeforeUsers{}
	credential, err := service.CreateClientAuthRalali()
	if err != nil {
		return result, err
	}

	data := &apiEntity.SsoRequest{}
	data.GrantType = os.Getenv("RALALI_GRANT_TYPES")
	data.ClientSecret = credential.ClientSecret
	data.ClientKey = credential.ClientKey

	jsonStr, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", os.Getenv("RALALI_OAUTH_URL")+"/v2/token", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return result, body
	}
	json.Unmarshal([]byte(body), &result)
	return result, nil
}

// AuthRalaliLogin params
// @token: string
// return ResultProfileOauth
func (service *V1SsoService) AuthRalaliLogin(payload *apiEntity.SsoAuthByPassword, waitGroup *sync.WaitGroup) (interface{}, error) {
	rllAauth, errData := service.GetRalaliOauthToken()
	if errData != nil {
		var messagesErr interface{}
		json.Unmarshal(errData, &messagesErr)
	}
	responseToken := &httpEntity.LoginOauthDetailResponse{}
	bodyRequest := apiEntity.SsoUserRequest{}
	bodyRequest.GrantType = "password"
	bodyRequest.Email = payload.Email
	bodyRequest.Password = payload.Password
	strRequest, _ := json.Marshal(bodyRequest)
	req, err := http.NewRequest("POST", os.Getenv("RALALI_OAUTH_URL")+"/v2/token", bytes.NewBuffer(strRequest))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+rllAauth.AccessToken)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	tokenUsers := apiEntity.SsoResponseAfterUsers{}
	err = json.Unmarshal([]byte(body), &tokenUsers)
	if err != nil {
		return tokenUsers, err
	}

	reqDataUser, err := http.NewRequest("GET", os.Getenv("RALALI_OAUTH_URL")+"/v2/profile", bytes.NewBuffer(strRequest))
	reqDataUser.Header.Set("Content-type", "application/json")
	reqDataUser.Header.Set("Authorization", "Bearer "+tokenUsers.AccessToken)

	if err != nil {
		return nil, err
	}
	clientUsers := &http.Client{}
	respUsers, err := clientUsers.Do(reqDataUser)
	// bodyUsers, _ := ioutil.ReadAll(respUsers.Body)
	// fmt.Println(string(bodyUsers))
	if err != nil {
		return nil, err
	}
	bodyUsers, _ := ioutil.ReadAll(respUsers.Body)

	userData := &apiEntity.ProfileOauthResponse{}
	err = json.Unmarshal([]byte(bodyUsers), &userData)
	if err != nil {
		return nil, err
	}
	var isRegistered bool

	rspDB, errstr := service.SsoBusiness.StoreRalaliUserProfile(userData)

	if errstr != nil {
		isRegistered = true
	}

	sessionDataUsers := httpEntity.SessionDataUsers{}
	sessionDataUsers.ID = rspDB.ID
	sessionDataUsers.Name = rspDB.Name
	sessionDataUsers.Email = rspDB.Email
	sessionDataUsers.IsOrganizers = rspDB.IsOrganizer
	expirationTime := time.Unix(1605194980, 0)

	sessdata, _ := json.Marshal(sessionDataUsers)
	claims := &httpEntity.Claims{}
	claims.SessionData = string(sessdata)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    rspDB.Email,
		Id:        string(rspDB.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return responseToken, err
	}
	responseToken.Token = tokenString
	responseToken.Expired = expirationTime.Unix()
	responseToken.IsRegistered = isRegistered
	responseToken.UserData.ID = rspDB.ID
	responseToken.UserData.Name = rspDB.Name
	responseToken.UserData.Email = rspDB.Email
	responseToken.UserData.Image = rspDB.Image
	responseToken.SessionData = sessionDataUsers
	return responseToken, nil
}
