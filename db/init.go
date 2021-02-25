package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() *sql.DB {
	driverName := "mysql"
	DsName := "root@(127.0.0.1:3306)/ca_dojo?charset=utf8"
	db, err := sql.Open(driverName, DsName)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// テーブルの作成
	var sql [3]string = [3]string{
		`CREATE TABLE IF NOT EXISTS user (
			id   INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name TEXT NOT NULL,
			token TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS userCharacter (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			userId INTEGER REFERENCES user(id),
			characterId INTEGER REFERENCES character(id)
		);`,
		`CREATE TABLE IF NOT EXISTS character (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			characterRank INTEGER NOT NULL,
			name TEXT NOT NULL
		);`,
	}

	for _, s := range sql {
		if _, err := db.Exec(s); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	println("Connected to the database")
	return db
}
