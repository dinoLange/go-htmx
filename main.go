package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	initDataMap()

	r := mux.NewRouter()

	r.HandleFunc("/{id}", makeSheetHandler())
	r.HandleFunc("/{field}/{id}", makeGetHandler()).Methods("GET")
	r.HandleFunc("/{field}/{id}", makePutHandler()).Methods("PUT")
	r.HandleFunc("/edit/{field}/{id}", makeEditHandler())

	log.Fatal(http.ListenAndServe(":8080", r))
}

func makeSheetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
			return
		}
		name, err := getNameById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderSheetTemplate(w, name)
	}
}

func makeGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		switch field {
		case "name":
			name, err := getNameById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderNameTemplate(w, name)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}
	}
}

func makePutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newName := r.FormValue("name")
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		switch field {
		case "name":
			name, err := putNameById(newName, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderNameTemplate(w, name)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}
	}
}

func makeEditHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "id "+vars["id"]+" is not an integer", http.StatusInternalServerError)
		}
		field := vars["field"]
		switch field {
		case "name":
			name, err := getNameById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			renderEditNameTemplate(w, name)
		default:
			http.Error(w, "editing field "+field+" not implemented", http.StatusInternalServerError)
		}

	}
}
