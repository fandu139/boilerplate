package mysql

import (
	"fmt"
	"net/url"
	"os"

	"github.com/sofyan48/boilerplate/src/utils/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbWrite *gorm.DB
	dbRead  *gorm.DB
)

// MySqlLibs ...
type MySqlLibs struct {
	Write *gorm.DB
	Read  *gorm.DB
}

// MySqlLibsHandler ...
func MySqlLibsHandler() *MySqlLibs {
	return &MySqlLibs{
		Write: writeDB(),
		Read:  readDB(),
	}
}

// DBInitTransaction Initialization Connection
// return connection, error
func dbInitTransaction(dbhost, dbport, dbuser, dbpass, dbname string) (*gorm.DB, error) {
	defaultTimezone := os.Getenv("SERVER_TIMEZONE")
	configDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname,
		url.QueryEscape(defaultTimezone),
	)
	DB, err := gorm.Open("mysql", configDB)
	defer DB.Close()
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return nil, err
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetConnMaxLifetime(10)
	DB.DB().SetMaxOpenConns(200)
	return DB, nil
}

// DBInitTransaction Initialization Connection
// return connection, error
func dbInitRead(dbhost, dbport, dbuser, dbpass, dbname string) (*gorm.DB, error) {
	defaultTimezone := os.Getenv("SERVER_TIMEZONE")
	configDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbuser,
		dbpass,
		dbhost,
		dbport,
		dbname,
		url.QueryEscape(defaultTimezone),
	)

	DB, err := gorm.Open("mysql", configDB)
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return nil, err
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetConnMaxLifetime(10)
	DB.DB().SetMaxOpenConns(200)
	return DB, nil
}

// readDB function
// return DB.Begin()
func readDB() *gorm.DB {
	dbhost := os.Getenv("DB_HOST_READ")
	dbport := os.Getenv("DB_PORT_READ")
	dbuser := os.Getenv("DB_USER_READ")
	dbpass := os.Getenv("DB_PASSWORD_READ")
	dbname := os.Getenv("DB_NAME_READ")
	if dbRead == nil {
		dbRead, err := dbInitRead(dbhost, dbport, dbuser, dbpass, dbname)
		if err != nil {
			fmt.Println(err)
			defer dbRead.Close()
		}
		return dbRead
	}
	return dbRead
}

// writeDB function
// return DB.Begin()
func writeDB() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if dbWrite == nil {
		dbWrite, err := dbInitTransaction(dbhost, dbport, dbuser, dbpass, dbname)
		if err != nil {
			fmt.Println(err)
			defer dbWrite.Close()
		}
		return dbWrite
	}
	return dbWrite
}

// func (libs *MySqlLibs) ShowConnection() {
// 	libs.Write.DB().Stats()
// }
