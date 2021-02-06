package record

import (
	"ca-tech-dojo/db"
	"encoding/base64"
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
func CreateUser(name string) string {
	db, _ := db.Connect()

	token := base64.StdEncoding.EncodeToString([]byte(name))

	const sql = "INSERT INTO user(name,token) values (?,?)"
	_, err := db.Exec(sql, name, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return token
}

// GetUser ユーザーの照会
func GetUser(token string) string {
	db, _ := db.Connect()

	// dbから取得
	const sql = "SELECT * FROM user WHERE token = ?"
	rows, err := db.Query(sql, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var u User
	for rows.Next() { //! この辺よくわからない
		// データをスキャン
		if err := rows.Scan(&u.ID, &u.Name, &u.Token); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	return u.Name
}

// UpdateUser ユーザー名の更新
func UpdateUser(newName string, token string) {
	db, _ := db.Connect()

	const sql = "UPDATE user SET name = ? WHERE token = ?;"
	_, err := db.Query(sql, newName, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
