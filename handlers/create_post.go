package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"forum/database"
)

// function to limit user spamming post creation
var userPostCreationTime = make(map[int]time.Time)
var userPostCreationCount = make(map[int]int)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}

	UserId, err := CheckAuthentication(w, r)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "))
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized UserProfile"))
		return
	}
	data := struct {
		Title          string
		Content        string
		Categories     []string
		CategoriesList []string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if (strings.Trim(data.Title, " ") == "") || (strings.Trim(data.Content, " ") == "") {
		ErrorJs(w, http.StatusBadRequest, errors.New("please enter title and content"))
		return
	}
	if (!title_RGX.MatchString(data.Title)) || (!content_RGX.MatchString(data.Content)) || (len(data.Categories) == 0) || (err != nil) {
		ErrorJs(w, http.StatusBadRequest, errors.New("required input not provided"))
		return
	}

	if time.Since(userPostCreationTime[UserId]) >= 5*time.Minute {
		userPostCreationCount[UserId] = 0
	}
	// Check if user has created too many posts in the given time frames
	if hasCreatedTooManyPostsIn5Minutes(UserId) {
		ErrorJs(w, http.StatusTooManyRequests, errors.New("too many posts created in a short period"))
		return
	}

	id, err := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("somethign went wrong creating post"))
		return
	}

	// Update user post creation time and count
	userPostCreationTime[UserId] = time.Now()
	userPostCreationCount[UserId]++
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "ok",
		"ID":            id,
		"Title":         data.Title,           // TODO get username
		"UserName":      UserProfile.UserName, // TODO get username
		"CreatedAt":     "now",
		"Content":       data.Content,
		"Categories":    data.CategoriesList,
		"LikeCount":     0,
		"DislikeCount":  0,
		"CommentsCount": 0,
	})
}

func hasCreatedTooManyPostsIn5Minutes(userId int) bool {
	if count, exists := userPostCreationCount[userId]; exists && count >= 5 {
		if time.Since(userPostCreationTime[userId]) <= 5*time.Minute {
			return true
		}
	}
	return false
}
