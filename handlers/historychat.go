package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetChatHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	userID, err := CheckAuthentication(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	otherUserID := r.URL.Query().Get("user_id")
	offSet := r.URL.Query().Get("offset")

	if otherUserID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	query := `
    SELECT 
        m.id,
        m.content,
        m.sender_id,
        m.receiver_id,
        m.timestamp,
        u.username
    FROM messages m
    JOIN users u ON m.sender_id = u.id
    WHERE 
        (m.sender_id = ? AND m.receiver_id = ?)
        OR 
        (m.sender_id = ? AND m.receiver_id = ?)
    ORDER BY m.timestamp DESC
    LIMIT 10 OFFSET ?`

	rows, err := DB.Query(query, userID, otherUserID, otherUserID, userID, offSet)
	if err != nil {
		fmt.Println("GetChatHistory", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(
			&msg.ID,
			&msg.Content,
			&msg.SenderID,
			&msg.ReceiverID,
			&msg.Timestamp,
			&msg.Username,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]Message{"messages": messages})
}
