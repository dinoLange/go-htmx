package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start Server")
	r := mux.NewRouter()

	r.HandleFunc("/", makeStartHandler())
	r.HandleFunc("/character", makeCreateHandler())
	r.HandleFunc("/character/{id}", makeSheetHandler()).Methods("GET")
	r.HandleFunc("/character/{id}", makeDeleteCharacterHandler()).Methods("DELETE")

	r.HandleFunc("/character/edit/basicattributes/{id}", makeEditBasicAttributesHandler()).Methods("GET")
	r.HandleFunc("/character/basicattributes/{id}", makePutBasicAttributesHandler()).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func makeStartHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		characters, err := loadAllCharacters()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		renderStartTemplate(w, characters)
	}
}

func makeCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		character := newCharacter()
		_, err := character.create()
		handleGenericError(err, w)
		fmt.Println("created: ", character)
		renderEditBasicAttributes(w, character)
	}
}

func newCharacter() Character {
	var character Character
	character.Age = 0
	character.Name = "name"
	character.Race = "race"
	character.Class = "class"
	character.Background = "bg"
	return character
}

func makeSheetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseId(r)
		handleGenericError(err, w)
		character, err := getCharacterById(id)
		handleGenericError(err, w)
		renderSheetTemplate(w, character)
	}
}

func makeDeleteCharacterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseId(r)
		handleGenericError(err, w)
		err = deleteCharacter(id)
		handleGenericError(err, w)
	}
}

func makeEditBasicAttributesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseId(r)
		handleGenericError(err, w)
		character, err := getCharacterById(id)
		handleGenericError(err, w)
		renderEditBasicAttributes(w, character)
	}
}

func makePutBasicAttributesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseId(r)
		handleGenericError(err, w)
		character := Character{Id: id}
		character.Name = r.FormValue("name")
		character.Age, err = strconv.ParseInt(r.FormValue("age"), 10, 64)
		if err != nil {
			http.Error(w, "age "+r.FormValue("age")+" is not an integer", http.StatusInternalServerError)
		}
		character.Race = r.FormValue("race")
		character.Class = r.FormValue("class")
		character.Background = r.FormValue("background")
		_, err = character.save()
		handleGenericError(err, w)
		renderBasicAttributes(w, character)
	}
}

func handleGenericError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseId(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	return strconv.ParseInt(vars["id"], 10, 64)
}
