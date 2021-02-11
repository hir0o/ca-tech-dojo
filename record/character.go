package record

import (
	"ca-tech-dojo/db"
	"fmt"
	"os"
)

type Character struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type DB struct {
	UserCharacterID  string
	UserID           int
	CharacterID      string
	UserCCharacterID int
	CharacterRank    int
	Name             string
}

func CharacterList(token string) []Character {
	db := db.Connect()

	// dbから取得
	const sql = "SELECT * FROM user WHERE token = ?"
	row := db.QueryRow(sql, token)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	const getUserCharacterID = "SELECT * FROM userCharacter WHERE user_id = ?"

	// userがもつcharacterを取得
	const getCharacterSql = "SELECT * FROM userCharacter INNER JOIN character ON character.id = userCharacter.characterId WHERE userCharacter.userId = ?;"
	rows, error := db.Query(getCharacterSql, u.ID)
	if error != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var characters []Character
	for rows.Next() {
		var d DB
		// 取得したデータを取得
		if err := rows.Scan(&d.UserCharacterID, &d.UserID, &d.CharacterID, &d.UserCCharacterID, &d.CharacterRank, &d.Name); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		characters = append(characters, Character{
			UserCharacterID: d.UserCharacterID,
			CharacterID:     d.CharacterID,
			Name:            d.Name,
		})
	}
	return characters
}
