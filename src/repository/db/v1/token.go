package v1

import (
	"encoding/json"
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	cache "github.com/ralali/event-api/src/util/helper/redis"
	"github.com/ralali/event-api/src/util/middleware"
)

// TokenRepository |
type TokenRepository struct {
	Cache             redis.Conn
	DefaultMiddleware middleware.DefaultMiddleware
}

// TokenRepositoryHandler handler all
// return : TokenRepository
func TokenRepositoryHandler() *TokenRepository {
	return &TokenRepository{
		Cache: cache.GetConnection(),
	}
}

// TokenRepositoryInterface interface
type TokenRepositoryInterface interface {
	GetSessionVisitor(context *gin.Context) (httpEntity.SessionDataUsers, error)
	GetSessionOrganizers(context *gin.Context) (httpEntity.SessionDataOrganizers, error)
}

// GetSessionVisitor params
// @token: string
// return interfac{}, error
func (repository *TokenRepository) GetSessionVisitor(context *gin.Context) (httpEntity.SessionDataUsers, error) {
	token := cache.GetToken(context)
	data := httpEntity.SessionDataUsers{}
	claims := &httpEntity.Claims{}
	repository.DefaultMiddleware.GetSessionClaim(token, claims)
	err := json.Unmarshal([]byte(claims.SessionData), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetSessionOrganizers params
// @token: string
// return interfac{}, error
func (repository *TokenRepository) GetSessionOrganizers(context *gin.Context) (httpEntity.SessionDataOrganizers, error) {
	token := cache.GetToken(context)
	data := &httpEntity.SessionDataUsers{}
	claims := &httpEntity.Claims{}
	results := httpEntity.SessionDataOrganizers{}
	repository.DefaultMiddleware.GetSessionClaim(token, claims)
	err := json.Unmarshal([]byte(claims.SessionData), &data)
	if err != nil {
		return results, err
	}
	if data.IsOrganizers != true {
		return results, errors.New("Your Not EO")
	}
	results = data.Organizers
	return results, nil
}
