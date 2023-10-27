package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db = initDataBaseAccess()

func initDataBaseAccess() *sql.DB {
	db, err := sql.Open("sqlite3", "data/main.db")
	if err != nil {
		log.Panic(err)
	}
	return db
}

func getCharacterById(id int64) (Character, error) {
	var character Character
	row := db.QueryRow("SELECT id, name FROM character WHERE id = ?", id)
	err := row.Scan(&character.Id, &character.Name)
	if err != nil {
		return character, err
	}
	return character, nil
}

func deleteCharacter(id int64) error {
	res, err := db.Exec("DELETE FROM character WHERE id = ?", id)

	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				return fmt.Errorf("Zero rows were deleted")
			}
			if count > 1 {
				return fmt.Errorf("More than one row was deleted")
			}
		}
	}
	return nil
}

func (character *Character) save() (int64, error) {
	fmt.Println("save char ", character)
	if _, err := getCharacterById(character.Id); err == sql.ErrNoRows {
		return 0, err
	}
	return character.update()
}

func (character *Character) create() (int64, error) {
	result, err := db.Exec("INSERT INTO character (name) VALUES (?)", character.Name)
	if err != nil {
		return 0, fmt.Errorf("create character failed: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("create character failed: %v", err)
	}
	character.Id = id
	return id, nil
}

func (character *Character) update() (int64, error) {
	fmt.Println(character)
	result, err := db.Exec("UPDATE character SET name = ? WHERE id = ?", character.Name, character.Id)
	if err != nil {
		return 0, fmt.Errorf("update character failed: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("update character failed: %v", err)
	}
	return id, nil
}

func loadAllCharacters() ([]Character, error) {
	var characters []Character
	rows, err := db.Query("SELECT id, name FROM character")
	if err != nil {
		return characters, err
	}
	var character Character
	for rows.Next() {
		rows.Scan(&character.Id, &character.Name)
		characters = append(characters, character)
	}
	return characters, nil

}
