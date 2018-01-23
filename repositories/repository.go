package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"edb/helpers"
)

var DBConn *sql.DB

type Repository struct {}

func (repository *Repository) getDB() (db *sql.DB) {
	if DBConn == nil {
		var err error

		DBConn, err = sql.Open(helpers.Config("db.driver"), helpers.Config("db.username")+":"+helpers.Config("db.password")+"@tcp("+helpers.Config("db.host")+":"+helpers.Config("db.port")+")/"+helpers.Config("db.database"))

		if err != nil {
			log.Fatal("Connect to database failed with error, ", err)
		}
	}

	return DBConn
}
