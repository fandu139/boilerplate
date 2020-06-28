package cockroach

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CockroachLibs ...
type CockroachLibs struct {
	Write *gorm.DB
	Read  *gorm.DB
}

// CockroachLibsHandler ...
func CockroachLibsHandler() *CockroachLibs {
	return &CockroachLibs{
		Write: getTransactionConnection(),
		Read:  getReadConnection(),
	}
}

// DBInitTransaction Initialization Connection
// return connection, error
func dbInit(dbhost, dbport, dbuser, dbpass, dbname string) (*gorm.DB, error) {
	var (
		configDB = fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			dbuser, dbpass, dbhost, dbport, dbname)
	)
	DB, err := gorm.Open("postgres", configDB)
	if err != nil {
		log.Println(fmt.Sprintf("failed to connect to database: %v", err))
		return nil, err
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetConnMaxLifetime(10)
	DB.DB().SetMaxOpenConns(200)
	return DB, nil
}

// GetTransactionConnection function
// return DB.Begin()
func getTransactionConnection() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	var TransactionDB *gorm.DB
	if TransactionDB == nil {
		TransactionDB, _ = dbInit(dbhost, dbport, dbuser, dbpass, dbname)
	}
	return TransactionDB
}

// getReadConnection function
// return DB.Begin()
func getReadConnection() *gorm.DB {
	dbhost := os.Getenv("DB_HOST_READ")
	dbport := os.Getenv("DB_PORT_READ")
	dbuser := os.Getenv("DB_USER_READ")
	dbpass := os.Getenv("DB_PASSWORD_READ")
	dbname := os.Getenv("DB_NAME_READ")
	var readConnection *gorm.DB
	if readConnection == nil {
		readConnection, _ = dbInit(dbhost, dbport, dbuser, dbpass, dbname)
	}
	return readConnection
}
