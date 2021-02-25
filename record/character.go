package record

import (
	"database/sql"
	"fmt"
	"os"
)

type Character struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

func CharacterList(token string, db *sql.DB) []Character {

	user, err := GetUser(token, db);
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	const getUserCharacterID = "SELECT * FROM usersCharacters WHERE user_id = ?"

	// userがもつcharacterを取得
	const getCharacterSQL = "SELECT * FROM own INNER JOIN usersCharacters ON usersCharacters.id = own.usersCharacterId WHERE own.userId = ?;"
	rows, error := db.Query(getCharacterSQL, user.ID)
	if error != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var characters []Character
	type CharacterDB struct {
		ID 									 int
		UsersCharacterID     string
		UserID               int
		CharacterID          string
		UserCharacterTableID int
		CharacterRank        int
		Name                 string
	}
	for rows.Next() {
		var c CharacterDB
		// 取得したデータを取得
		if err := rows.Scan(&c.ID, &c.UserID, &c.UsersCharacterID, &c.UserCharacterTableID, &c.CharacterID, &c.CharacterRank, &c.Name); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		characters = append(characters, Character{
			UserCharacterID: c.UsersCharacterID,
			CharacterID:     c.CharacterID,
			Name:            c.Name,
		})
	}
	return characters
}
