package database

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, email, username, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare("INSERT INTO users(email ,username, password) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(email, username, hashedPassword)
	if err != nil {
		return 0, err
	}
	uid, err := res.LastInsertId()
	if err != nil {
		return int(uid), err
	}
	fmt.Println("User created successfully!")
	return int(uid), nil
}

func GetUserByUname(db *sql.DB, username string) (string, int, error) {
	var hpassword string
	var uid int
	err := db.QueryRow("SELECT id,password FROM users WHERE username=? collate NOCASE", username).Scan(&uid, &hpassword)

	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return "", 0, fmt.Errorf("user not found")
	} else if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	return hpassword, uid, nil
}

func GetUserByUemail(db *sql.DB, email string) (string, int, error) {
	var hpassword string
	var uid int
	err := db.QueryRow("SELECT id,password FROM users WHERE email=? collate NOCASE", email).Scan(&uid, &hpassword)

	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return "", 0, fmt.Errorf("user not found invalid password or username")
	} else if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	return hpassword, uid, nil
}

func AddSessionToken(db *sql.DB, user_id int, token string) error {
	if token == "" || user_id == 0 {
		return fmt.Errorf("invalid token or user_id")
	}
	stmt, err := db.Prepare("INSERT INTO sessions(user_id, session_token, expiration) VALUES(?,?,DATETIME(CURRENT_TIMESTAMP, '+1 hour'))")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, token)
	if err != nil {
		return err
	}
	fmt.Println("Session token created successfully!")
	return nil
}

// Returns 0 if uid not found
func GetUidFromToken(db *sql.DB, token string) (int, error) {
	if token == "" {
		// case of new user
		return 0, nil
	}
	var uid int
	err := db.QueryRow("SELECT user_id FROM sessions WHERE session_token=?", token).Scan(&uid)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("session not found")
	} else if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uid, nil
}
