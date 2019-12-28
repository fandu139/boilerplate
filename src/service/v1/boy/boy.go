package boy

import (
	"github.com/jinzhu/gorm"
	connection "github.com/sofyan48/BOILERGOLANG/src/util/helper/mysqlconnection"
)

// V1Boys | Derivated from UserRepository
type V1Boys struct {
	DB gorm.DB
	// Redis redis.Conn
}
type Boys struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// V1BoysHandler Handler
func V1BoysHandler() *V1Boys {
	return &V1Boys{
		DB: *connection.GetConnection(),
		// Redis: redisConn.GetConnection(),
	}
}

//V1BoysInterface declare All Method
type V1BoysInterface interface {
	Boys() *Boys
}

// Boys Function
// return HealthResponse
func (service *V1Boys) Boys() *Boys {
	result := &Boys{}
	result.Name = "Vhaldi"
	result.Status = "Menyesal Karena Yanda"
	return result
}
