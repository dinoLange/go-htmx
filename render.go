package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Character struct {
	Name string
}

func renderSheetTemplate(w http.ResponseWriter, name string) {
	character := Character{Name: name}

	err := templates.ExecuteTemplate(w, "sheet.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderNameTemplate(w http.ResponseWriter, name string) {
	character := Character{Name: name}
	err := templates.ExecuteTemplate(w, "name.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderEditNameTemplate(w http.ResponseWriter, name string) {
	character := Character{Name: name}
	err := templates.ExecuteTemplate(w, "edit-name.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}