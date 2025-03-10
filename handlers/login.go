package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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
	JsResponse(w, 200, true, nil)
	// http.Redirect(w, r, "/", http.StatusFound)
}

func MiddleWear(w http.ResponseWriter, r *http.Request) (bool, int) {
	userID, err := CheckAuthentication(w, r)
	fmt.Println("User id", userID, "err", err)
	if userID == 0 || err != nil {
		return false, 0
	}
	return true, userID
}

func JsResponse(w http.ResponseWriter, status int, msgStatus bool, data any) {
	w.Header().Set("Content-Type", "application-json")
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application-json")
	}
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	json.NewEncoder(w).Encode(map[string]any{
		"status": msgStatus,
		"data":   data,
	})
}

func Checker(w http.ResponseWriter, r *http.Request) {
	check, UID := MiddleWear(w, r)
	status := 200
	if !check {
		status = http.StatusUnauthorized
	}
	userName, err := database.GetUsernameByUid(DB, UID)
	if err != nil {
		check = false
		status = http.StatusUnauthorized
	}
	var data map[string]interface{}
	if check {
		data = map[string]any{
			"UserName": userName,
		}
	}
	JsResponse(w, status, check, data)
}
func Login(w http.ResponseWriter, r *http.Request) {
	var data = map[string]any{
		"status":  false,
		"token":   "",
		"message": "",
	}
	fmt.Println("Start Connect")
	Data := struct {
		Email    string
		Password string
	}{}
	json.NewDecoder(r.Body).Decode(&Data)
	fmt.Println("=========================\n", Data)
	r.ParseForm()
	if Data.Email == "" || Data.Password == "" {
		data["message"] = "Email or username is required"
		JsResponse(w, http.StatusUnauthorized, false, data)
		return
	}
	if !validpassword(Data.Password) {
		data["message"] = "Invalid Password"
		JsResponse(w, http.StatusUnauthorized, false, data)
		return
	}
	fmt.Println("Invalid Password")

	var hpassword string
	var uid int
	var err error
	hpassword, uid, err = database.GetUserByUemail(DB, Data.Email)
	if err != nil || hpassword == "" {
		hpassword, uid, err = database.GetUserByUname(DB, Data.Email)
		if err != nil || hpassword == "" {
			data["message"] = "Invalid Email or Password"
			JsResponse(w, http.StatusUnauthorized, false, data)
			return
		}
	}
	fmt.Println("Invalid email or username")

	err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(Data.Password))
	if err != nil {
		data["message"] = "Invalid Email or Usermame"
		JsResponse(w, http.StatusUnauthorized, false, data)
		return
	}

	fmt.Println("sdf")

	token, err := tokening.GenerateSessionToken("email:" + Data.Email)
	if err != nil {
		data["message"] = "Something wrong please try later..."
		JsResponse(w, http.StatusInternalServerError, false, data)
		return
	}
	fmt.Println("tdg")

	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		data["message"] = "Something wrong please try again later..."
		JsResponse(w, http.StatusInternalServerError, false, data)
		return
	}
	SetCookie(w, token, "session", true)
	fmt.Println("Connected Sucssefully")
	// func JsResponse(w http.ResponseWriter, status int, msgStatus bool, data interface{}) {
	data["status"] = true
	data["token"] = token
	JsResponse(w, http.StatusOK, true, data)

	// JsResponse(w, true, "User Logged Succesfully")
	// Commented Part Bellow is belong to the old Forum PAGE :D

	// redirected := RedirectToHomeIfAuthenticated(w, r)
	// if redirected {
	// 	return
	// }
	// r.ParseForm()
	// email := strings.ToLower(r.Form.Get("email"))
	// upass := r.Form.Get("password")
	// structError := map[string]interface{}{
	// 	"StatuCode":    http.StatusBadRequest,
	// 	"MessageError": errors.New("username or email is required"),
	// 	"Register":     false,
	// }
	// if email == "" || upass == "" {
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }
	// if !validpassword(upass) {
	// 	structError["MessageError"] = "invalid password"
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }
	// var hpassword string
	// var uid int
	// var err error
	// hpassword, uid, err = database.GetUserByUemail(DB, email)
	// if err != nil || hpassword == "" {
	// 	hpassword, uid, err = database.GetUserByUname(DB, email)
	// 	if err != nil || hpassword == "" {
	// 		structError["StatuCode"] = http.StatusBadRequest
	// 		structError["MessageError"] = "invalid email or username"
	// 		structError["Register"] = false
	// 		ErrorPage(w, "register.html", structError)
	// 		return
	// 	}
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(upass))
	// if err != nil {
	// 	structError["StatuCode"] = http.StatusUnauthorized
	// 	structError["MessageError"] = "invalid email or password"
	// 	structError["Register"] = false
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }

	// token, err := tokening.GenerateSessionToken("email:" + email)
	// if err != nil {
	// 	structError["StatuCode"] = http.StatusInternalServerError
	// 	structError["MessageError"] = "error creating a session"
	// 	structError["Register"] = false
	// 	ErrorPage(w, "error.html", structError)
	// 	return
	// }
	// err = database.AddSessionToken(DB, uid, token)
	// if err != nil {
	// 	log.Println(err)
	// 	structError["StatuCode"] = http.StatusInternalServerError
	// 	structError["MessageError"] = "Internal server error"
	// 	structError["Register"] = false
	// 	ErrorPage(w, "error.html", structError)
	// 	return
	// }
	// SetCookie(w, token, "session", true)
	// http.Redirect(w, r, "/", http.StatusFound)
}

func Register(w http.ResponseWriter, r *http.Request) {

	var Data struct {
		Email     string
		UserName  string
		Gender    string
		FirstName string
		LastName  string
		Password  string
		Age       string
	}
	var RespondData = map[string]any{
		"status":  false,
		"token":   "",
		"message": "",
	}
	json.NewDecoder(r.Body).Decode(&Data)
	// redirected := RedirectToHomeIfAuthenticated(w, r)
	// if redirected {
	// 	return
	// }
	// r.ParseForm()
	// uemail := r.Form.Get("email")
	// uname := r.Form.Get("username")
	// upass := r.Form.Get("password")
	// structError := map[string]interface{}{
	// 	"StatuCode":    http.StatusBadRequest,
	// 	"MessageError": errors.New("email or username already taken"),
	// 	"Register":     true,
	// }
	// Gender    string
	// 	FirstName string
	// 	LastName  string
	if (!email_RGX.MatchString(Data.Email)) ||
		(!username_RGX.MatchString(Data.UserName)) {
		RespondData["message"] = "Email or Username is not valid."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	// r.ParseForm()

	// // Get form values
	// firstName := r.Form.Get("first-name")
	// lastName := r.Form.Get("last-name")
	// gender := r.Form.Get("gender")
	// email := r.Form.Get("email")
	// username := r.Form.Get("username")
	// password := r.Form.Get("password")

	// age, err := strconv.Atoi(r.Form.Get("age"))

	// structError := map[string]interface{}{
	// 	"StatuCode":    http.StatusBadRequest,
	// 	"MessageError": errors.New("invalid input"),
	// 	"Register":     true,
	// }

	// Validate required fields
	// if firstName == "" || lastName == "" || gender == "" || age <= 0 {
	// 	structError["MessageError"] = "all fields are required"
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }

	// // Validate email, username and password
	// if !email_RGX.MatchString(email) || !username_RGX.MatchString(username) || !validpassword(password) {
	// 	structError["MessageError"] = "invalid email or username"
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }

	exist := CheckUserExists(Data.Email, Data.UserName)
	if exist {
		RespondData["message"] = "email or username already taken."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}

	if !(validpassword(Data.Password)) {
		RespondData["message"] = "Password is not valid."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	if strings.Trim(Data.FirstName, " ") == "" && strings.Trim(Data.LastName, " ") == "" {
		RespondData["message"] = "Please Enter your First And Last Name."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	Data.Gender = strings.ToLower(Data.Gender)
	if Data.Gender != "male" && Data.Gender != "female" {
		RespondData["message"] = "Please Enter your Gender."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	fmt.Println("Age", Data.Age)
	Age, err := strconv.Atoi(Data.Age)
	if err != nil {
		RespondData["message"] = "Please Enter the age."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	if Age < 18 {
		RespondData["message"] = "You are under the age... Try again after couple years."
		JsResponse(w, http.StatusBadRequest, false, RespondData)
		return
	}
	if Age > 140 {
		RespondData["message"] = "its seems you entered incorrect age.."
		JsResponse(w, http.StatusUnauthorized, false, RespondData)
		return
	}

	// Create user
	uid, err := database.CreateUser(DB, Data.FirstName, Data.LastName, Data.Gender, Age, Data.Email, Data.UserName, Data.Password)
	if err != nil {
		RespondData["message"] = "something wrong, please try later"
		JsResponse(w, http.StatusInternalServerError, false, RespondData)
		return
	}
	// exist := CheckUserExists(uemail, uname)
	// if exist {
	// 	structError["MessageError"] = "email or username already taken"
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }

	// uid, err = database.CreateUser(DB, Data.Email, Data.UserName, Data.Password)
	// if err != nil {
	// 	structError["StatuCode"] = http.StatusInternalServerError
	// 	structError["MessageError"] = "something wrong, please try later"
	// 	ErrorPage(w, "register.html", structError)
	// 	return
	// }
	token, err := tokening.GenerateSessionToken("username:" + Data.UserName)
	if err != nil {
		RespondData["message"] = "Something went wrong, try again later.."
		JsResponse(w, http.StatusInternalServerError, false, RespondData)
		return
	}

	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		RespondData["message"] = "Something went wrong, try again later.."
		JsResponse(w, http.StatusInternalServerError, false, RespondData)
		return
	}

	SetCookie(w, token, "session", true)
	RespondData["status"] = true
	RespondData["token"] = token
	RespondData["message"] = "User Created succesfully"
	JsResponse(w, http.StatusInternalServerError, true, RespondData)
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
