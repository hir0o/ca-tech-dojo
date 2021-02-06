package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// dbと接続するための関数
func Connect() *sql.DB {
	driverName := "mysql"
	DsName := "root@(127.0.0.1:3306)/ca_dojo?charset=utf8"
	db, err := sql.Open(driverName, DsName)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return db
}

// migration
func Init() {
	db := Connect()

	// テーブルの作成
	var sql [1]string = [1]string{ // TODO: 他のテーブルに関しても追加する
		`CREATE TABLE IF NOT EXISTS user (
			id   INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name TEXT NOT NULL,
			token TEXT NOT NULL
		);`,
	}

	for _, s := range sql {
		if _, err := db.Exec(s); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	println("Connected to the database")
}
