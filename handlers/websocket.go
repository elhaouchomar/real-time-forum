// handlers/chat.go
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	wsLib "github.com/gorilla/websocket"
)

var (
	upgrader = wsLib.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Configure appropriately for production
		},
	}

	// Thread-safe connected users map
	connectedUsers = struct {
		sync.RWMutex
		m map[int]*UserConnection
	}{m: make(map[int]*UserConnection)}
)

type UserConnection struct {
	Conn     *wsLib.Conn
	Username string
	UserID   int
}

type Message struct {
	ID         int       `json:"id,omitempty"`
	Content    string    `json:"content"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Timestamp  time.Time `json:"timestamp"`
	Type       string    `json:"type"`
	Username   string    `json:"username"`
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID, err := CheckAuthentication(w, r)
	if err != nil {
		// http.Er/ror(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var username string
	err = DB.QueryRow("SELECT username FROM users WHERE id = ?;", userID).Scan(&username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	userConn := &UserConnection{
		Conn:     conn,
		Username: username,
		UserID:   userID,
	}

	connectedUsers.Lock()
	connectedUsers.m[userID] = userConn
	connectedUsers.Unlock()

	broadcastStatus(userID, username, true)

	defer func() {
		conn.Close()
		connectedUsers.Lock()
		delete(connectedUsers.m, userID)
		connectedUsers.Unlock()
		broadcastStatus(userID, username, false)
	}()

	// Handle incoming messages
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		fmt.Println(msg)
		if err != nil {
			if !wsLib.IsCloseError(err, wsLib.CloseGoingAway, wsLib.CloseAbnormalClosure) {
				log.Printf("WebSocket read errorfff: %v", err)
			}
			break
		}
		var receiver_id int
		err = DB.QueryRow("SELECT id FROM users WHERE username = ?;", msg.Username).Scan(&receiver_id)
		 if err != nil   || receiver_id == userID {
			// http.Error(w, "User not found", http.StatusNotFound)
			// return
			// delet user from map
			connectedUsers.Lock()
			delete(connectedUsers.m, userID)
			connectedUsers.Unlock()
			break

		 }
		msg.SenderID = userID
		msg.ReceiverID = receiver_id
		msg.Timestamp = time.Now()

		// Save message to database
		err = saveMessage(msg)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}

		// Send message
		if msg.ReceiverID == 0 {
			broadcastMessage(msg)
		} else {
			sendPrivateMessage(msg)
		}
	}
}

func saveMessage(msg Message) error {
	query := `INSERT INTO messages (content, sender_id, receiver_id, timestamp) 
              VALUES (?, ?, ?, ?)`
	result, err := DB.Exec(query, msg.Content, msg.SenderID, msg.ReceiverID, msg.Timestamp)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	msg.ID = int(id)
	return nil
}

func GetChatHistory(w http.ResponseWriter, r *http.Request) {
	userID, err := CheckAuthentication(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	otherUserID := r.URL.Query().Get("user_id")
	if otherUserID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	query := `
        SELECT m.id, m.content, m.sender_id, m.receiver_id, m.timestamp, u.username
        FROM messages m
        JOIN users u ON m.sender_id = u.id
        WHERE (m.sender_id = ? AND m.receiver_id = ?) 
           OR (m.sender_id = ? AND m.receiver_id = ?)
        ORDER BY m.timestamp DESC
        LIMIT 50`

	rows, err := DB.Query(query, userID, otherUserID, otherUserID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID,
			&msg.Timestamp, &msg.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}

	json.NewEncoder(w).Encode(messages)
}

func broadcastStatus(userID int, username string, online bool) {
	status := Message{
		Type:      "status",
		Content:   username,
		SenderID:  userID,
		Timestamp: time.Now(),
	}
	if online {
		status.Content += " is online"
	} else {
		status.Content += " is offline"
	}
	broadcastMessage(status)
}

func broadcastMessage(msg Message) {
	connectedUsers.RLock()
	defer connectedUsers.RUnlock()

	for _, conn := range connectedUsers.m {
		if conn.UserID != msg.SenderID {
			err := conn.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message to user %d: %v", conn.UserID, err)
			}
		}
	}
}

func sendPrivateMessage(msg Message) {
	connectedUsers.Lock()
	defer connectedUsers.Unlock()

	fmt.Println(connectedUsers, msg)
	if conn, ok := connectedUsers.m[msg.ReceiverID]; ok {
		err := conn.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error sending private message: %v", err)
		}
	} else {
		log.Printf("User %d is not connected", msg.ReceiverID)
	}
}
