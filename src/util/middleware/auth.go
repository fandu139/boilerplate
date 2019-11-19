package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// AuthToken fucntion
// return gin.HandlerFunc
func (m *DefaultMiddleware) AuthToken(levels int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if levels == 1 {
			tokenString := m.GetToken(c)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRET_KEY")), nil
			})

			if token == nil && err != nil {
				c.Next()
			}
		} else {
			tokenString := m.GetToken(c)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRET_KEY")), nil
			})
			if token != nil && err == nil {
				c.Next()
			} else {
				result := gin.H{
					"message": "not authorized",
					"error":   err.Error(),
				}
				c.JSON(http.StatusUnauthorized, result)
				c.Abort()
			}
		}
	}
}

// StringWithCharset params
// @length: int
// return string
func (m *DefaultMiddleware) StringWithCharset(length int) string {
	b := make([]byte, length)
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetToken params
// @context: *gin.Context
// return gin.HandlerFunc
func (m *DefaultMiddleware) GetToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"]
	if len(token) < 1 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token Not Compatible",
		})
		context.Abort()
		return "event"
	}
	return token[0]
}

// GetSessionClaim params
func (m *DefaultMiddleware) GetSessionClaim(tokenString string, claims *httpEntity.Claims) *httpEntity.Claims {
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return claims, nil
	})
	// fmt.Println(claims)
	return claims
}
