package record

import (
	"ca-tech-dojo/db"
	"ca-tech-dojo/lib"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type CharacterDB struct {
	ID            string
	CharacterRank int
	Name          string
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

func GachaDraw(times int, token string) []GachaResult {
	db := db.Connect()
	var characters []CharacterDB

	gachaTimes := lib.GenerateWeightedNumber(times)
	for i, t := range gachaTimes { // rankごとに、キャラクターを取得
		if t == 0 { // 回数が0だったらbreak
			continue
		}
		// rankと、数を指定して取得
		const sql = "SELECT * FROM characters WHERE (characters.characterRank = ?) ORDER BY RAND() LIMIT ?;"
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
			characters = append(characters, c)
		}
	}

	// userを取得
	var user User
	user, _ = GetUser(token)

	// 結果を格納する変数
	var gachaResults []GachaResult

	// 取得したcharactorをusersCharacterと、ownテーブルに保存
	for _, character := range characters {
		const usersCharactersSQL = "INSERT INTO usersCharacters(characterRank,characterName) values (?,?)"
		r, err := db.Exec(usersCharactersSQL, character.CharacterRank, character.Name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		const ownSQL = "INSERT INTO own(userId,usersCharacterId) values (?,?)"
		if _, err := db.Exec(ownSQL, user.ID, id); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		gachaResults = append(gachaResults, GachaResult{
			CharacterID: character.ID,
			Name: character.Name,
		})
	}
	return gachaResults
}
