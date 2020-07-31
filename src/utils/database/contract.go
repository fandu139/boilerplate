package database

import (
	"github.com/sofyan48/boilerplate/src/utils/database/cockroach"
	"github.com/sofyan48/boilerplate/src/utils/database/mysql"
	"github.com/sofyan48/boilerplate/src/utils/database/postgre"
)

type Contract interface {
	Postgre() *postgre.PostgreLibs
	MySQL() *mysql.MySqlLibs
	Roach() *cockroach.CockroachLibs
}
