// handlers/chat.go
package handlers

import (
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

func GetUsers(userID int) ([]string, error) {
	var users []string
	rows, err := DB.Query("SELECT username FROM users WHERE id != ?;", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		users = append(users, username)
	}
	return users, nil
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID, err := CheckAuthentication(w, r)
	if err != nil {
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
	BroadcastUsersList()
	GetChatHistory(userID)
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
		if err != nil || receiver_id == userID {

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
		sendPrivateMessage(msg)

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

type ChatHistory struct {
	Type     string    `json:"type"`
	Messages []Message `json:"messages"`
}

func GetChatHistory(userID int) {
	query := `
        SELECT m.id, m.content, m.sender_id, m.receiver_id, m.timestamp, u.username
        FROM messages m
        JOIN users u ON m.sender_id = u.id
        WHERE m.sender_id = ? OR m.receiver_id = ?
        ORDER BY m.timestamp ASC`

	rows, err := DB.Query(query, userID, userID)
	if err != nil {
		log.Printf("Error retrieving chat history: %v", err)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID,
			&msg.Timestamp, &msg.Username)
		if err != nil {
			log.Printf("Error scanning message: %v", err)
			return
		}
		msg.Type = "chat_history"
		messages = append(messages, msg)
	}

	history := ChatHistory{
		Type:     "chat_history",
		Messages: messages,
	}

	connectedUsers.RLock()
	if conn, ok := connectedUsers.m[userID]; ok {
		err := conn.Conn.WriteJSON(history)
		if err != nil {
			log.Printf("Error sending chat history: %v", err)
		}
	}
	connectedUsers.RUnlock()
}
func broadcastStatus(userID int, username string, online bool) {
	status := Message{
		Type:      "status",
		Content:   username,
		SenderID:  userID,
		Timestamp: time.Now(),
	}
	if online {
		status.Content = "online"
	} else {
		status.Content = "offline"
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

type UsersList struct {
	Type         string            `json:"type"`
	Usernames    []string          `json:"usernames"`
	UserStatuses map[string]string `json:"user_statuses"`
	UserIDs      map[string]int    `json:"user_ids"`
}

func BroadcastUsersList() error {
	connectedUsers.Lock()
	defer connectedUsers.Unlock()

	for currentUserID, conn := range connectedUsers.m {
		usernames, err := GetUsers(currentUserID)
		if err != nil {
			log.Printf("Error getting users for user %d: %v", currentUserID, err)
			continue
		}

		userStatuses := make(map[string]string)
		userIDs := make(map[string]int)

		// Populate user data
		for _, username := range usernames {
			// Get user ID
			userID, err := GetUserIDByUsername(username)
			if err != nil {
				log.Printf("Error getting user ID for %s: %v", username, err)
				continue
			}
			userIDs[username] = userID

			// Set status
			userStatuses[username] = "offline"
			for _, c := range connectedUsers.m {
				if c.Username == username {
					userStatuses[username] = "online"
					break
				}
			}
		}

		usersList := UsersList{
			Type:         "users_list",
			Usernames:    usernames,
			UserStatuses: userStatuses,
			UserIDs:      userIDs,
		}

		err = conn.Conn.WriteJSON(usersList)
		if err != nil {
			log.Printf("Error sending users list to user %d: %v", conn.UserID, err)
		}
	}
	return nil
}

func GetUserIDByUsername(username string) (int, error) {
	var id int
	err := DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	return id, err
}
func sendPrivateMessage(msg Message) {
	connectedUsers.Lock()
	defer connectedUsers.Unlock()

	if conn, ok := connectedUsers.m[msg.ReceiverID]; ok {
		err := conn.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error sending private message: %v", err)
		}
	} else {
		log.Printf("User %d is not connected", msg.ReceiverID)
	}
}
