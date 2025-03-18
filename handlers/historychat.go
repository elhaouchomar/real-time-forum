package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetChatHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Headers:", r.Header)
	fmt.Println("Request URL with Domain:", r.Host+r.URL.String())
	fmt.Println("Request Cookies:", r.Cookies())
	check, userID := MiddleWear(w, r)
	if !check {
		ErrorJs(w, http.StatusUnauthorized, errors.New("not authorized"))
		return
	}

	otherUserID := r.URL.Query().Get("user_id")
	offSet := r.URL.Query().Get("offset")

	if otherUserID == "" {
		ErrorJs(w, http.StatusBadRequest, errors.New("missing user_id parameter"))
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
		ErrorJs(w, http.StatusInternalServerError, errors.New("error apply query"))
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
			ErrorJs(w, http.StatusInternalServerError, errors.New("internale server error "))
			return
		}
		messages = append(messages, msg)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]Message{"messages": messages})
}
