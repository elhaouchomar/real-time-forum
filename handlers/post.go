package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/database"
	"forum/structs"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	template := getHtmlTemplate()

	var Postid int

	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	Postid, err = strconv.Atoi(r.PathValue("id"))
	if err != nil || Postid < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, err := database.GetPostByID(DB, Postid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comments, err := database.GetCommentsByPost(DB, uid, post.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	template.ExecuteTemplate(w, "post.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}

func InfiniteScroll(w http.ResponseWriter, r *http.Request) {
	check, uid := MiddleWear(w, r)
	if !check || uid == 0 {
		JsResponse(w, 401, false, nil)
		return
	}

	var profile structs.Profile
	var err error
	profile, err = database.GetUserProfile(DB, uid)
	if err != nil {
		log.Fatal(err)
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
			username = profile.UserName
		}
		profile, err = database.GetUserProfile(DB, username)

		if err != nil {
			ErrorJs(w, http.StatusNotFound, errors.New("page not found"))
			return
		}
		posts, err = database.QuerryPostsbyUser(DB, username, uid, structs.Limit, offset)
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
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(struct {
		Posts      []structs.Post  `json:"posts"`
		Profile    structs.Profile `json:"profile"`
		Categories map[string]int  `json:"categories"`
	}{Posts: posts, Profile: profile, Categories: categories})
	fmt.Println(err)
}
