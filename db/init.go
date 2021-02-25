package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// migration
func Init() *sql.DB{
	driverName := "mysql"
	DsName := "root@(127.0.0.1:3306)/ca_dojo?charset=utf8"
	db, err := sql.Open(driverName, DsName)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// テーブルの作成
	var sql [4]string = [4]string{
		`CREATE TABLE IF NOT EXISTS users (
			id    INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name  TEXT NOT NULL,
			token TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS own (
			id               INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			userId           INTEGER NOT NULL,
			usersCharacterId INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS usersCharacters (
			id            INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			characterId TEXT NOT NULL,
			characterRank INTEGER NOT NULL,
			characterName TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS characters (
			id            INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			characterRank INTEGER NOT NULL,
			name          TEXT NOT NULL
		);`,
	}

	for _, s := range sql {
		if _, err := db.Exec(s); err != nil {
			panic(err)
		}
	}
	println("Connected to the database")

	return db;
}
