package storage

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB
var gormDB *gorm.DB

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@(localhost:3305)/realworld?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = gormDB.Exec("select 1").Error
	if err != nil {
		panic(err)
	}
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
