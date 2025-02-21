package database

import (
	"database/sql"
	"errors"
	"forum/structs"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) CreateMessage(content string, authorID int64) error {
	query := `INSERT INTO messages (content, author_id) VALUES (?, ?)`
	_, err := db.Exec(query, content, authorID)
	return err
}

func (db *DB) GetMessages() ([]structs.Message, error) {
	query := `
        SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
        FROM messages m 
        JOIN users u ON m.author_id = u.id 
        ORDER BY m.timestamp DESC
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []structs.Message
	for rows.Next() {
		var msg structs.Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.AuthorID, &msg.Author, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (db *DB) GetMessagesByUser(userID int64) ([]structs.Message, error) {
	query := `
        SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
        FROM messages m 
        JOIN users u ON m.author_id = u.id 
        WHERE m.author_id = ?
        ORDER BY m.timestamp DESC
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []structs.Message
	for rows.Next() {
		var msg structs.Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.AuthorID, &msg.Author, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (db *DB) GetMessageByID(msgID int64) (*structs.Message, error) {
	query := `
        SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
        FROM messages m 
        JOIN users u ON m.author_id = u.id 
        WHERE m.id = ?
    `
	var msg structs.Message
	err := db.QueryRow(query, msgID).Scan(
		&msg.ID, &msg.Content, &msg.AuthorID, &msg.Author, &msg.Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("message not found")
	}
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (db *DB) DeleteMessage(msgID, userID int64) error {
	query := `DELETE FROM messages WHERE id = ? AND author_id = ?`
	result, err := db.Exec(query, msgID, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("message not found or unauthorized")
	}
	return nil
}

func (db *DB) GetChatHistory(user1ID, user2ID int64) ([]structs.Message, error) {
	query := `
        SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
        FROM messages m 
        JOIN users u ON m.author_id = u.id 
        WHERE (m.author_id = ? AND m.recipient_id = ?) 
           OR (m.author_id = ? AND m.recipient_id = ?)
        ORDER BY m.timestamp ASC
    `
	rows, err := db.Query(query, user1ID, user2ID, user2ID, user1ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []structs.Message
	for rows.Next() {
		var msg structs.Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.AuthorID, &msg.Author, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}



func (db *DB) GetMessage() []structs.Message {
    query := `
        SELECT m.id, m.content, m.author_id, u.username, m.timestamp 
        FROM messages m 
        JOIN users u ON m.author_id = u.id 
        ORDER BY m.timestamp DESC
    `
    rows, err := db.Query(query)
    if err != nil {
        return nil
    }
    defer rows.Close()

    var messages []structs.Message
    for rows.Next() {
        var msg structs.Message
        err := rows.Scan(&msg.ID, &msg.Content, &msg.AuthorID, &msg.Author, &msg.Timestamp)
        if err != nil {
            continue
        }
        messages = append(messages, msg)
    }
    return messages
}