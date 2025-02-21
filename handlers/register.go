package handlers

import (
	"net/http"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	redirected := RedirectToHomeIfAuthenticated(w, r)
	if redirected {
		return
	}
	getHtmlTemplate().ExecuteTemplate(w, "register.html",
		struct {
			Register     bool
			MessageError string
		}{Register: r.URL.Path == "/register", MessageError: ""})
}
