package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const USER = "name"
const PASS = "pass"
const DBNAME = "database"

func Connect() *sql.DB {
	URL := fmt.Sprint("%s:%s@tcp(127.0.0.1:3306)/%s", USER, PASS, DBNAME)
	con, erro := sql.Open(DRIVER, URL)
	if erro != nil {
		panic(erro.Error())
	}
	return con
}
