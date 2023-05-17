package main

import (
	"fmt"
	"os"
)

func main() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_USER_PWD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// デバッグ用に値を出力
	fmt.Printf("mysqlUser: %s\nmysqlUserPwd: %s\nmysqlDatabase: %s\n", mysqlUser, mysqlUserPwd, mysqlDatabase)

}
