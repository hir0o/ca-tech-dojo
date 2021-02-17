package record

import (
	"ca-tech-dojo/db"
	"ca-tech-dojo/lib"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Charactor struct {
	characterID int
	name        string
}

type CharacterDB struct {
	ID            int
	CharactorRank int
	Name          string
}

type GachaCharacter struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

func GachaDraw(times int, token string) []GachaCharacter {
	db := db.Connect()
	var characters []GachaCharacter

	gachaTimes := lib.GenerateWeightedNumber(times)
	for i, t := range gachaTimes { // rankごとに、キャラクターを取得
		if t == 0 { // 回数が0だったらbreak
			continue
		}
		// rankと、数を指定して取得
		const sql = "SELECT * FROM character WHERE (character.characterRank = ?) ORDER BY RAND() LIMIT ?;"
		rows, err := db.Query(sql, i, t)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		for rows.Next() {
			var c CharacterDB
			// 取得したデータを取得
			if err := rows.Scan(&c.ID, &c.CharacterRank, &c.Name); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			// 引いたキャラクターを保存
			charactors = append(charactors, Charactor{c.ID, c.Name})
		}
	}

	// userを取得
	const getUserSQL = "SELECT * FROM user WHERE token = ?"
	row := db.QueryRow(getUserSQL, token)

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Token); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 取得したcharactorをuserCharactorテーブルに保存
	for _, charactor := range charactors {
		const sql = "INSERT INTO userCharactor(userId,charactorId) values (?,?)"
		_, err := db.Exec(sql, u.ID, charactor.characterID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return characters
}
