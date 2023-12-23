package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func init() {
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3305)/realword")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
