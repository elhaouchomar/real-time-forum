package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/database"
	"forum/structs"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}

	var Postid int
	// TODO check for edge cases
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusUnauthorized,
			"MessageError": "unauthorized " + err.Error(),
		})
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusUnauthorized,
			"MessageError": "unauthorized " + err.Error(),
		})
		return
	}
	Postid, err = strconv.Atoi(r.PathValue("id"))
	if err != nil || Postid < 0 {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusBadRequest,
			"MessageError": "Badrequest invalid post id " + err.Error(),
		})
		return
	}
	post, err := database.GetPostByID(DB, Postid, uid)
	if err != nil {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusInternalServerError,
			"MessageError": "internal server error" + err.Error(),
		})
		return
	}

	comments, err := database.GetCommentsByPost(DB, uid, post.ID)
	if err != nil {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusInternalServerError,
			"MessageError": "internal server error trying to get comments " + err.Error(),
		})
	}

	template.ExecuteTemplate(w, "post.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}

func InfiniteScroll(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "))
		return
	}

	var profile structs.Profile
	if uid != 0 {
		fmt.Println("Getting Profile Information...")
		profile, err = database.GetUserProfile(DB, uid)
		if err != nil {
			log.Fatal(err)
		}
	}
	typeQuery := r.URL.Query().Get("type")
	offset_str := r.URL.Query().Get("offset")
	fmt.Println("typeQuery", typeQuery, "offset_str", offset_str)
	fmt.Println("typeQuery", r.URL.Query().Get("type"), "offset_str", r.URL.Query().Get("offset"))
	offset, err := strconv.Atoi(offset_str)
	if err != nil {
		offset = 0
	}

	var posts []structs.Post
	switch typeQuery {
	case "category":
		category := r.URL.Query().Get("category")
		if !database.IsCategoryValid(category) {
			ErrorJs(w, http.StatusBadRequest, errors.New("invalid category"))
			return
		}
		posts, err = database.QuerryLatestPostsByCategory(DB, uid, category, offset)
		fmt.Println("Posts :", posts)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	case "profile":
		username := r.URL.Query().Get("username")
		if username == "" {
			fmt.Println("--------5----------", username)
			username = profile.UserName
		}
		profile, err = database.GetUserProfile(DB, username)

		if err != nil {
			ErrorJs(w, http.StatusNotFound, errors.New("page not found"))
			return
		}
		posts, err = database.QuerryPostsbyUser(DB, username, uid, structs.Limit, offset)
		// if len(posts) > 0 {
		// 	posts[0].UserPostsCounts = profile.ArticleCount
		// 	posts[0].UsersCommentsCounts = profile.CommentCount
		// }
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, errors.New("error fetching posts "))
			return
		}
	case "home":
		posts, err = database.QuerryLatestPosts(DB, uid, structs.Limit, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	case "liked":
		posts, err = database.QuerryLatestPostsByUserLikes(DB, uid, structs.Limit, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	case "trending":
		posts, err = database.QuerryMostLikedPosts(DB, uid, structs.Limit, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	default:
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid url"))
		return
	}

	categories, err := database.GetCategoriesWithPostCount(DB)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}
	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Optionally set the status code to 200 OK
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(struct {
		Posts      []structs.Post  `json:"posts"`
		Profile    structs.Profile `json:"profile"`
		Categories map[string]int  `json:"categories"`
	}{Posts: posts, Profile: profile, Categories: categories})
	fmt.Println(err)
}
