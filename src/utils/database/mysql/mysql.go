package mysql

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		log.Fatal(fmt.Sprintf("failed to connect to database: %v", err))
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
		log.Println(fmt.Sprintf("failed to connect to database: %v", err))
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
	var readDB *gorm.DB
	if readDB == nil {
		readDB, err := dbInitRead(dbhost, dbport, dbuser, dbpass, dbname)
		if err != nil {
			fmt.Println(err)
			defer readDB.Close()
		}
		return readDB
	}
	return nil
}

// writeDB function
// return DB.Begin()
func writeDB() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	var TransactionDB *gorm.DB
	if TransactionDB == nil {
		TransactionDB, err := dbInitTransaction(dbhost, dbport, dbuser, dbpass, dbname)
		if err != nil {
			fmt.Println(err)
			defer TransactionDB.Close()
		}
		return TransactionDB
	}
	return nil
}

// func (libs *MySqlLibs) ShowConnection() {
// 	libs.Write.DB().Stats()
// }
