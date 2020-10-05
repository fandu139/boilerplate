package database

import (
	"github.com/fandu139/boilerplate/src/utils/database/cockroach"
	"github.com/fandu139/boilerplate/src/utils/database/mysql"
	"github.com/fandu139/boilerplate/src/utils/database/postgre"
)

type Contract interface {
	Postgre() *postgre.PostgreLibs
	MySQL() *mysql.MySqlLibs
	Roach() *cockroach.CockroachLibs
}
