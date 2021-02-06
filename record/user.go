package record

import (
	"ca-tech-dojo/db"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// データ型
type User struct {
	ID    int64
	Name  string
	Token string
}

// データの作成
func CreateUser(name string) string {
	db := db.Connect()
	token := "123456789" // TODO: tokenを生成する
	const sql = "INSERT INTO user(name,token) values (?,?)"
	_, err := db.Exec(sql, name, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return token
}
