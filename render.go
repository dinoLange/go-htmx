package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderStartTemplate(w http.ResponseWriter, characters []Character) {
	err := templates.ExecuteTemplate(w, "start.html", characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderSheetTemplate(w http.ResponseWriter, character Character) {
	err := templates.ExecuteTemplate(w, "sheet.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderNameTemplate(w http.ResponseWriter, character Character) {
	err := templates.ExecuteTemplate(w, "name.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderEditNameTemplate(w http.ResponseWriter, character Character) {
	err := templates.ExecuteTemplate(w, "edit-name.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
