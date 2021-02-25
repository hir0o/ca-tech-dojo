package record

import (
	"ca-tech-dojo/db"
	"ca-tech-dojo/lib"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// User データ型
type User struct {
	ID    int64
	Name  string
	Token string
}

// CreateUser データの作成
func CreateUser(name string) (string, error) {
	db := db.Connect()

	token := lib.GenerateRandomString()

	const sql = "INSERT INTO users(name,token) values (?,?)"
	_, err := db.Exec(sql, name, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return token, nil
}

// GetUser ユーザーの照会
func GetUser(token string) (User, error) {
	db := db.Connect()

	// dbから取得
	const sql = "SELECT * FROM users WHERE token = ?"
	row := db.QueryRow(sql, token)

	var u User
	// データをスキャン
	err := row.Scan(&u.ID, &u.Name, &u.Token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return u, err
}

// UpdateUser ユーザー名の更新
func UpdateUser(newName string, token string) error {
	db := db.Connect()

	const sql = "UPDATE users SET name = ? WHERE token = ?;"
	_, err := db.Exec(sql, newName, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return err
}
