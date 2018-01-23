package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"edb/helpers"
)

var DBConn *sql.DB
var TxConn *sql.Tx

type Repository struct {}

func (repository *Repository) getTx() *sql.Tx {
	if TxConn != nil {
		return TxConn
	}

	var err error

	db := repository.getDB()
	TxConn, err = db.Begin()

	if err != nil {
		log.Fatal("Connect to database failed with error, ", err)
	}

	return TxConn
}

func (repository *Repository) getDB() *sql.DB {
	if DBConn == nil {
		var err error

		DBConn, err = sql.Open(helpers.Config("db.driver"), helpers.Config("db.username")+":"+helpers.Config("db.password")+"@tcp("+helpers.Config("db.host")+":"+helpers.Config("db.port")+")/"+helpers.Config("db.database"))

		if err != nil {
			log.Fatal("Connect to database failed with error, ", err)
		}
	}

	return DBConn
}
