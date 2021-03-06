package record

import (
	"ca-tech-dojo/lib"
	"database/sql"
	"fmt"
	"math/rand"
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

func GachaDraw(times int, token string, db *sql.DB) ([]GachaResult, error) {
	var characters []CharacterDB

	gachaTimes := lib.GenerateWeightedNumber(times)
	for i, t := range gachaTimes { // rankごとに、キャラクターを取得
		if t == 0 { // 回数が0だったらbreak
			continue
		}
		for j := 0; j < t; j++ {
			// rankと、数を指定して取得
			const sql = "SELECT * FROM characters WHERE (characters.characterRank = ?) ORDER BY RAND() LIMIT 1;"
			row := db.QueryRow(sql, i)

			var c CharacterDB
			// 取得したデータを取得
			if err := row.Scan(&c.ID, &c.CharacterRank, &c.Name); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil, err
			}
			// 引いたキャラクターを保存
			characters = append(characters, c)
		}
	}

	// キャラクターの順番をシャッフル
	rand.Shuffle(len(characters), func(i, j int) {
		characters[i], characters[j] = characters[j], characters[i]
	})

	// userを取得
	var user User
	user, _ = GetUser(token, db)
	// 結果を格納する変数
	var gachaResults []GachaResult

	// 取得したcharactorをusersCharacterと、ownテーブルに保存
	for _, character := range characters {
		const usersCharactersSQL = "INSERT INTO usersCharacters(characterId,characterRank,characterName) values (?,?,?)"
		r, err := db.Exec(usersCharactersSQL, character.ID, character.CharacterRank, character.Name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		const ownSQL = "INSERT INTO own(userId,usersCharacterId) values (?,?)"
		if _, err := db.Exec(ownSQL, user.ID, id); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		gachaResults = append(gachaResults, GachaResult{
			CharacterID: character.ID,
			Name: character.Name,
		})
	}
	return gachaResults, nil
}
