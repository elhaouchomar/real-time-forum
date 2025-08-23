package handlers

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	template := getHtmlTemplate()
	template.ExecuteTemplate(w, "index.html", nil)
}
