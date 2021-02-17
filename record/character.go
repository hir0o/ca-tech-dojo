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

func CharacterList(token string) []Character {
	db := db.Connect()

	// dbから取得
	const sql = "SELECT * FROM users WHERE token = ?"
	row := db.QueryRow(sql, token)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	const getUserCharacterID = "SELECT * FROM usersCharacters WHERE user_id = ?"

	// userがもつcharacterを取得
	const getCharacterSQL = "SELECT * FROM usersCharacters INNER JOIN characters ON characters.id = usersCharacters.characterId WHERE usersCharacters.userId = ?;"
	rows, error := db.Query(getCharacterSQL, u.ID)
	if error != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var characters []Character
	type CharacterDB struct {
		UserCharacterID      string
		UserID               int
		CharacterID          string
		UserCharacterTableID int
		CharacterRank        int
		Name                 string
	}
	for rows.Next() {
		var c CharacterDB
		// 取得したデータを取得
		if err := rows.Scan(&c.UserCharacterID, &c.UserID, &c.CharacterID, &c.UserCharacterTableID, &c.CharacterRank, &c.Name); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		characters = append(characters, Character{
			UserCharacterID: c.UserCharacterID,
			CharacterID:     c.CharacterID,
			Name:            c.Name,
		})
	}
	return characters
}
