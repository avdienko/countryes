package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "127.0.0.1"
	port     = "3306"
	user     = "root"
	password = "123456"
	dbName   = "db"
)

var instance *sql.DB

func GetInstance() (*sql.DB, error) {
	if instance == nil {
		db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbName)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		instance = db
	}
	return instance, nil
}
