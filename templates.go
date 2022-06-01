package main

import (
	"html/template"
	"net/http"
	"regexp"
)

const (
	editTemplate = "./templates/edit.html"
	viewTemplate = "./templates/view.html"
)

var templates = template.Must(template.ParseFiles(editTemplate, viewTemplate))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
