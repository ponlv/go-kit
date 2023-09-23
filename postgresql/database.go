package postgresql

import (
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func GetDB() *pg.DB {
	return db
}
