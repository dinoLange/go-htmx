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

func getCharacterById(id int) (Character, error) {
	var character Character
	row := db.QueryRow("SELECT id, name FROM character WHERE id = ?", id)
	err := row.Scan(&character.Id, &character.Name)
	if err == sql.ErrNoRows {
		_, err := character.create()
		if err != nil {
			return character, err
		}
	}
	if err != nil {
		return character, err
	}
	return character, nil
}

func (character *Character) save() (int64, error) {
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
	return id, nil
}

func (character *Character) update() (int64, error) {
	result, err := db.Exec("UPDATE character SET name = ?", character.Name)
	if err != nil {
		return 0, fmt.Errorf("update character failed: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("update character failed: %v", err)
	}
	return id, nil
}
