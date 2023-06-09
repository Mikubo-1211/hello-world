package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// ①-1
	// ①-1
	//mysqlUser := "test_user"
	//mysqlUserPwd := "password"
	//mysqlDatabase := "test_database"

	// ①-2
	//_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, _ := sql.Open("mysql", connStr)
	// ①-2

	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}
