package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB {
	driver, t_string := AccessString()
	conn, err := sql.Open(driver, t_string)
	// conn, err := sql.Open("mysql", "maath:_14326@tcp(127.0.0.1:3306)/db_go_react")

	if err != nil {
		panic(err.Error())
	}
	return conn
}

func DbQuery(conn *sql.DB, query string) *sql.Rows {
	result, err := conn.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
	return result
}
