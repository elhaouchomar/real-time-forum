package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"forum/database"
)

func PostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	UserId, err := CheckAuthentication(w, r)
	if err != nil || UserId == 0 {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "))
		return
	}

	var requestData struct {
		PostID string `json:"postId"`
		Type   string `json:"type"`
		Post   bool   `json:"post"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid jsonX"))
		return
	}
	if requestData.Type != "like" && requestData.Type != "dislike" {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid type"))
		return
	}
	PostIdInt, err := strconv.Atoi(requestData.PostID)
	if err != nil {
		ErrorJs(w, 400, errors.New("invalid like post id"))
		return
	}
	if PostIdInt < 0 {
		ErrorJs(w, 400, errors.New("invalid like post id"))
	}
	liked, err := database.HasUserLikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}

	dislike, err := database.HasUserDislikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}
	addeddStatus := false
	if requestData.Type == "like" {
		if dislike {
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if liked {
			// remove the like from the post in database
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.LikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	} else {
		if liked {
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if dislike {
			// remove the like from the post in database
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.DislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				ErrorJs(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	}

	// get the new like count
	likeCount, err := database.GetPostLikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting like count"))
		return
	}
	dislikeCount, err := database.GetPostDislikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting like count"))
		return
	}
	database.UpdatePostLikeCount(DB, PostIdInt, requestData.Post)
	database.UpdatePostDislikeCount(DB, PostIdInt, requestData.Post)
	// return the new like count
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "ok",
		"likes":    likeCount,
		"dislikes": dislikeCount,
		"added":    addeddStatus,
	})
}
