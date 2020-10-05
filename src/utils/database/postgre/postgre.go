package postgre

import (
	"fmt"
	"os"

	"github.com/fandu139/boilerplate/src/utils/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	dbWrite *gorm.DB
	dbRead  *gorm.DB
)

// PostgreLibs ...
type PostgreLibs struct {
	Write *gorm.DB
	Read  *gorm.DB
}

// PostgreLibsHandler ...
func PostgreLibsHandler() *PostgreLibs {
	return &PostgreLibs{
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
		log.Fatal("failed to connect to database", err)
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
	if dbWrite == nil {
		dbWrite, _ = dbInit(dbhost, dbport, dbuser, dbpass, dbname)
	}
	return dbWrite
}

// getReadConnection function
// return DB.Begin()
func getReadConnection() *gorm.DB {
	dbhost := os.Getenv("DB_HOST_READ")
	dbport := os.Getenv("DB_PORT_READ")
	dbuser := os.Getenv("DB_USER_READ")
	dbpass := os.Getenv("DB_PASSWORD_READ")
	dbname := os.Getenv("DB_NAME_READ")
	if dbRead == nil {
		dbRead, _ = dbInit(dbhost, dbport, dbuser, dbpass, dbname)
	}
	return dbRead
}
