package handlers

import (
	"errors"
	"forum/database"
	"net/http"
)

var structError = map[string]interface{}{
	"StatuCode":    http.StatusBadRequest,
	"MessageError": errors.New("email or username already taken"),
	"Register":     false,
}

func CheckAuthentication(w http.ResponseWriter, r *http.Request) (userID int, err error) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		structError["StatuCode"] = http.StatusUnauthorized
		structError["MessageError"] = "unauthorized " + err.Error()
		ErrorPage(w, "error.html", structError)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	userID, err = database.GetUidFromToken(DB, c.Value)
	if err != nil {
		Logout(w, r)
		return
	}
	return
}

func RedirectToHomeIfAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	userID, err := CheckAuthentication(w, r)
	if userID != 0 && err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	return false
}

func CheckUserExists(uemail, uname string) bool {
	check := 0
	hashpassrd, userId, err := database.GetUserByUname(DB, uname)
	if (hashpassrd != "" && userId != 0) || err == nil {
		check++
	}
	hashpassrd, userId, err = database.GetUserByUemail(DB, uemail)
	if (hashpassrd != "" && userId != 0) || err == nil {
		check++
	}
	return check > 0
}
