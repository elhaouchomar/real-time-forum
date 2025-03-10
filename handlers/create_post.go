package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"forum/database"
)

// function to limit user spamming post creation
var userPostCreationTime = make(map[int]time.Time)
var userPostCreationCount = make(map[int]int)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var DATA = map[string]any{
		"status":  false,
		"message": "",
	}
	if r.Method != "POST" {
		DATA["message"] = "invalid method"
		JsResponse(w, http.StatusMethodNotAllowed, false, DATA)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		DATA["message"] = "invalid content type"
		JsResponse(w, http.StatusBadRequest, false, DATA)
		return
	}

	UserId, err := CheckAuthentication(w, r)
	if err != nil {
		DATA["message"] = "unauthorized..."
		JsResponse(w, http.StatusUnauthorized, false, DATA)
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)
	if err != nil {
		DATA["message"] = "Unauthorized UserProfile..."
		JsResponse(w, http.StatusUnauthorized, false, DATA)
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
		DATA["message"] = "Please check title and content..."
		JsResponse(w, http.StatusBadRequest, false, DATA)
		return
	}
	if (!title_RGX.MatchString(data.Title)) || (!content_RGX.MatchString(data.Content)) || (len(data.Categories) == 0) || (err != nil) {
		DATA["message"] = "required input not provided..."
		JsResponse(w, http.StatusBadRequest, false, DATA)
		return
	}

	if time.Since(userPostCreationTime[UserId]) >= 5*time.Minute {
		userPostCreationCount[UserId] = 0
	}
	// Check if user has created too many posts in the given time frames
	if hasCreatedTooManyPostsIn5Minutes(UserId) {
		DATA["message"] = "You have reached the limit of creating posts, wait for 5 minutes"
		JsResponse(w, http.StatusTooManyRequests, false, DATA)
		return
	}

	id, err := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		DATA["message"] = "Somethign went wrong, please try again later..."
		JsResponse(w, http.StatusInternalServerError, false, DATA)
		return
	}

	// Update user post creation time and count
	userPostCreationTime[UserId] = time.Now()
	userPostCreationCount[UserId]++
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        true,
		"ID":            id,
		"Title":         data.Title,
		"UserName":      UserProfile.UserName,
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
