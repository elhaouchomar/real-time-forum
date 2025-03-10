package handlers

import (
	"encoding/json"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, page string, data map[string]interface{}) {
	w.WriteHeader(data["StatuCode"].(int))
	getHtmlTemplate().ExecuteTemplate(w, page, data)
}

func ErrorJs(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
