package handlers

import (
	"html/template"
	"log"
)

func getHtmlTemplate() *template.Template {
	template, err := template.New("index").ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	return template
}
