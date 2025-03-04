package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"forum/database"
	tokening "forum/handlers/token"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB                                              *sql.DB
	email_RGX, username_RGX, title_RGX, content_RGX *regexp.Regexp
)

func init() {

	email_RGX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	username_RGX = regexp.MustCompile(`^\w{3,19}$`)
	title_RGX = regexp.MustCompile(`.{1,60}`)
	content_RGX = regexp.MustCompile(`.{1,512}`)
}

func Error(w http.ResponseWriter, r *http.Request) {
	template := getHtmlTemplate()

	r.ParseForm()
	code, err := strconv.Atoi(r.Form.Get("code"))
	message := r.Form.Get("message")

	if err != nil || code < 100 || code > 599 {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusBadRequest,
			"MessageError": errors.New("invalid status code"),
		})
		return
	}

	w.WriteHeader(code)
	template.ExecuteTemplate(w, "error.html", map[string]interface{}{
		"StatuCode":    code,
		"MessageError": message,
	})
}
func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteAllCookie(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func MiddleWear(w http.ResponseWriter, r *http.Request) bool {
	userID, err := CheckAuthentication(w, r)
	fmt.Println("User id", userID, "err", err)
	if userID == 0 || err != nil {
		userID = 0
		return false
	}
	return true
}

func JsResponse(w http.ResponseWriter, status int, msgStatus bool, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": msgStatus,
		"data":   data,
	})
}

func Checker(w http.ResponseWriter, r *http.Request) {
	check := MiddleWear(w, r)
	status := 200
	if !check {
		status = http.StatusUnauthorized
	}
	var data map[string]interface{}
	if check {
		data = map[string]interface{}{
			"Test": "Checker Function in golang",
		}
	}
	JsResponse(w, status, check, data)
}
func Login(w http.ResponseWriter, r *http.Request) {
	redirected := RedirectToHomeIfAuthenticated(w, r)
	if redirected {
		return
	}
	r.ParseForm()
	email := strings.ToLower(r.Form.Get("email"))
	upass := r.Form.Get("password")
	structError := map[string]interface{}{
		"StatuCode":    http.StatusBadRequest,
		"MessageError": errors.New("username or email is required"),
		"Register":     false,
	}
	if email == "" || upass == "" {
		ErrorPage(w, "register.html", structError)
		return
	}
	if !validpassword(upass) {
		structError["MessageError"] = "invalid password"
		ErrorPage(w, "register.html", structError)
		return
	}
	var hpassword string
	var uid int
	var err error
	hpassword, uid, err = database.GetUserByUemail(DB, email)
	if err != nil || hpassword == "" {
		hpassword, uid, err = database.GetUserByUname(DB, email)
		if err != nil || hpassword == "" {
			structError["StatuCode"] = http.StatusBadRequest
			structError["MessageError"] = "invalid email or username"
			structError["Register"] = false
			ErrorPage(w, "register.html", structError)
			return
		}
	}
	fmt.Println(uid,email,upass)
	err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(upass))
	if err != nil {
		structError["StatuCode"] = http.StatusUnauthorized
		structError["MessageError"] = "invalid email or password"
		structError["Register"] = false
		ErrorPage(w, "register.html", structError)
		return
	}

	token, err := tokening.GenerateSessionToken("email:" + email)
	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "error creating a session"
		structError["Register"] = false
		ErrorPage(w, "error.html", structError)
		return
	}
	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		log.Println(err)
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "Internal server error"
		structError["Register"] = false
		ErrorPage(w, "error.html", structError)
		return
	}
	SetCookie(w, token, "session", true)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Register(w http.ResponseWriter, r *http.Request) {
	redirected := RedirectToHomeIfAuthenticated(w, r)
	if redirected {
		return
	}
	r.ParseForm()

	firstName := r.Form.Get("first-name")
	lastName := r.Form.Get("last-name")
	gender := r.Form.Get("gender")
	email := r.Form.Get("email")
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	age, err := strconv.Atoi(r.Form.Get("age"))

	structError := map[string]interface{}{
		"StatuCode":    http.StatusBadRequest,
		"MessageError": errors.New("invalid input"),
		"Register":     true,
	}

	if firstName == "" || lastName == "" || gender == "" || age <= 0 {
		structError["MessageError"] = "all fields are required"
		ErrorPage(w, "register.html", structError)
		return
	}

	if !email_RGX.MatchString(email) || !username_RGX.MatchString(username) || !validpassword(password) {
		structError["MessageError"] = "invalid email or username"
		ErrorPage(w, "register.html", structError)
		return
	}

	exist := CheckUserExists(email, username)
	if exist {
		structError["MessageError"] = "email or username already taken"
		ErrorPage(w, "register.html", structError)
		return
	}

	uid, err := database.CreateUser(DB, firstName, lastName, gender, age, email, username, password)
	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "something wrong, please try later"
		ErrorPage(w, "register.html", structError)
		return
	}

	token, err := tokening.GenerateSessionToken("username:" + username)
	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "something wrong, please try later"
		ErrorPage(w, "register.html", structError)
		return
	}

	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		structError["StatuCode"] = http.StatusInternalServerError
		structError["MessageError"] = "something wrong, please try later"
		ErrorPage(w, "register.html", structError)
		return
	}

	SetCookie(w, token, "session", true)
	http.Redirect(w, r, "/", http.StatusFound)
}

func validpassword(password string) bool {
	// Lowercase UPPERCASE digit {symbol}
	var a, A, d, s bool
	if len(password) < 8 || len(password) > 64 {
		return false
	}
	for _, char := range password {
		if !a && unicode.IsLower(char) {
			a = true
			continue
		} else if !A && unicode.IsUpper(char) {
			A = true
			continue
		} else if !d && unicode.IsDigit(char) {
			d = true
			continue
		} else if !s && !unicode.IsDigit(char) && !unicode.IsLetter(char) {
			s = true
			continue
		}
		if a && A && d && s {
			return true
		}
	}
	return a && A && d && s
}
