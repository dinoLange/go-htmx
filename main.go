package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Character struct {
	Id         int64
	Name       string
	Age        int64
	Race       string
	Class      string
	Background string
}

func main() {
	fmt.Println("Start Server")
	r := mux.NewRouter()

	r.HandleFunc("/", makeStartHandler())
	r.HandleFunc("/character", makeCreateHandler())
	r.HandleFunc("/character/{id}", makeSheetHandler()).Methods("GET")
	r.HandleFunc("/character/{id}", makeDeleteCharacterHandler()).Methods("DELETE")
	r.HandleFunc("/{field}/{id}", makeGetHandler()).Methods("GET")
	r.HandleFunc("/{field}/{id}", makePutHandler()).Methods("PUT")
	r.HandleFunc("/character/basicattributes/{id}", makePutBasicAttributesHandler()).Methods("PUT")
	r.HandleFunc("/edit/{field}/{id}", makeEditHandler())

	log.Fatal(http.ListenAndServe(":8080", r))
}

func makeCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var character Character
		_, err := character.create()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("created: ", character)
		renderEditBasicAttributes(w, character)
	}
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

func makeSheetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
			return
		}
		character, err := getCharacterById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderSheetTemplate(w, character)
	}
}

func makeDeleteCharacterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
			return
		}
		err = deleteCharacter(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func makeGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		switch field {
		case "name":
			character, err := getCharacterById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderBasicAttributes(w, character)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}
	}
}

func makePutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newName := r.FormValue("name")
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		character := Character{Id: id}
		switch field {
		case "name":
			character.Name = newName
			_, err := character.save()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderBasicAttributes(w, character)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}
	}
}

func makePutBasicAttributesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		character := Character{Id: id}
		character.Name = r.FormValue("name")
		character.Age, err = strconv.ParseInt(r.FormValue("age"), 10, 64)
		if err != nil {
			http.Error(w, "age "+r.FormValue("age")+" is not an integer", http.StatusInternalServerError)
		}
		character.Race = r.FormValue("race")
		character.Class = r.FormValue("class")
		_, err = character.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		renderBasicAttributes(w, character)
	}
}

func makeEditHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		switch field {
		case "name":
			character, err := getCharacterById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderEditBasicAttributes(w, character)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}

	}
}
