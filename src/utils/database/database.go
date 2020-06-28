package database

import (
	"github.com/sofyan48/boilerplate/src/utils/database/cockroach"
	"github.com/sofyan48/boilerplate/src/utils/database/mysql"
	"github.com/sofyan48/boilerplate/src/utils/database/postgre"
)

type Database struct{}

func DatabaseHandler() *Database {
	return &Database{}
}

type DatabaseInterface interface {
	Postgre() *postgre.PostgreLibs
	MySQL() *mysql.MySqlLibs
	Roach() *cockroach.CockroachLibs
}

// Postgre ..
func (db *Database) Postgre() *postgre.PostgreLibs {
	return postgre.PostgreLibsHandler()
}

// MySQL ...
func (db *Database) MySQL() *mysql.MySqlLibs {
	return mysql.MySqlLibsHandler()
}

// Roach ...
func (db *Database) Roach() *cockroach.CockroachLibs {
	return cockroach.CockroachLibsHandler()
}
