package handlers

import (
	"fmt"
	"forum/structs"
	"log"
	"net/http"

	"forum/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		structError["StatuCode"] = http.StatusNotFound
		structError["MessageError"] = "page not found"
		ErrorPage(w, "error.html", structError)
		return
	}
	template := getHtmlTemplate()

	userId, err := CheckAuthentication(w, r)
	
	if err != nil {
		return
	}

	var profile structs.Profile
	if userId != 0 {
		profile, err = database.GetUserProfile(DB, userId)
		if err != nil {
			log.Fatal(err)
		}
	}

	categories, err := database.GetCategoriesWithPostCount(DB)

	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "error getting categories from database " + err.Error()
		ErrorPage(w, "error.html", structError)
		return
	}
	if r.FormValue("type") != "" {
		profile.CurrentPage = r.FormValue("type")
		profile.Category = r.FormValue("category")
	} else {
		profile.CurrentPage = "home"
	}
	categor := structs.Categories(categories)

	err = template.ExecuteTemplate(w, "index.html", struct {
		Posts      []structs.Post
		Profile    structs.Profile
		Categories structs.Categories
	}{nil, profile, categor})
	if err != nil {
		fmt.Println("error executing template", err)
	}
}
