package main

import (
	"fmt"
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

func renderBasicAttributes(w http.ResponseWriter, character Character) {
	err := templates.ExecuteTemplate(w, "basic-attributes.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderEditBasicAttributes(w http.ResponseWriter, character Character) {
	fmt.Println(character)
	err := templates.ExecuteTemplate(w, "edit-basic-attributes.html", character)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
