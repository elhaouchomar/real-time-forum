package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"forum/database"
	"forum/structs"
)

var userCommmentCreationTime = make(map[int]time.Time)

// Create a map to store user comment creation count
var userCommmentCreationCount = make(map[int]int)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	session, err := r.Cookie("session")
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	UserId, err := database.GetUidFromToken(DB, session.Value)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	data := struct {
		PostID  string
		Comment string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if len(data.Comment) == 0 || strings.TrimSpace(data.Comment) == "" {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid comment"))
		return
	}
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	IdInt, err := strconv.Atoi(data.PostID)
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid post id"))
		return
	}
	if time.Since(userCommmentCreationTime[UserId]) >= 1*time.Minute {
		userCommmentCreationCount[UserId] = 0
	}
	if hasCreatedTooManyCommentsIn5Minutes(UserId) {
		ErrorJs(w, http.StatusTooManyRequests, errors.New("too many posts created in a short period"))
		return
	}
	err, id := database.CreateComment(DB, UserId, IdInt, data.Comment)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating comment"))
		return
	}
	var post structs.Post
	post, err = database.GetPostByID(DB, IdInt, UserId)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting post"))
		return
	}
	userCommmentCreationTime[UserId] = time.Now()

	userCommmentCreationCount[UserId]++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       "ok",
		"UserName":     UserProfile.UserName,
		"CreatedAt":    "just now",
		"CommentID":    id,
		"Content":      data.Comment,
		"CommentCount": post.CommentCount,
	})
}

func hasCreatedTooManyCommentsIn5Minutes(userId int) bool {
	if count, exists := userCommmentCreationCount[userId]; exists && count >= 10 {
		if time.Since(userCommmentCreationTime[userId]) <= 1*time.Minute {
			return true
		}
	}
	return false
}
