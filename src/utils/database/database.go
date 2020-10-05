package database

import (
	"github.com/fandu139/boilerplate/src/utils/database/cockroach"
	"github.com/fandu139/boilerplate/src/utils/database/mysql"
	"github.com/fandu139/boilerplate/src/utils/database/postgre"
)

type database struct{}

// New ..
func New() Contract {
	return &database{}
}

// Postgre ..
func (db *database) Postgre() *postgre.PostgreLibs {
	return postgre.PostgreLibsHandler()
}

// MySQL ...
func (db *database) MySQL() *mysql.MySqlLibs {
	return mysql.MySqlLibsHandler()
}

// Roach ...
func (db *database) Roach() *cockroach.CockroachLibs {
	return cockroach.CockroachLibsHandler()
}
