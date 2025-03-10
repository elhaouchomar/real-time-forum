package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(file string) *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database is actually accessible
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database opened successfully!")

	// Execute PRAGMA statements
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		log.Fatalf("Error setting PRAGMA foreign_keys: %v", err)
	}

	return db
}

func CreateTables(db *sql.DB) {
	for t, c := range tables {
		_, err := db.Exec(c)
		if err != nil {
			log.Fatalf("Error creating table %s: %v", t, err)
		}
		fmt.Printf("Created table: %s\n", t)
	}
	DeleteExpiredSessions(db)
}


func CreateTriggers(db *sql.DB) {
	for _, c := range trigers {
		if len(c.tables) == 0 {
			rc, err := db.Exec(c.statment)
			if err != nil && err.Error() != "trigger "+c.name+" already exists" {
				log.Fatalf("Error creating TRIGGER %v: %v, '%v'\n", c.name, err, rc)
				continue
			}
			fmt.Printf("Created trigger1: %s\n", c.name)
		} else {
			for _, todo := range c.tables {
				azer := strings.ReplaceAll(c.statment, "1here2", todo)
				// fmt.Println(azer)
				res, err := db.Exec(azer)
				if err != nil && err.Error() != "trigger "+todo+c.name+" already exists" {
					log.Fatalf("Error creating TRIGGER %v: %v, '%v'\n%v\n", todo+c.name, err, todo, res)
					continue
				}
				fmt.Printf("Created trigger2: %s\n", todo+c.name)
			}
		}
	}
}

func DES_Ticker(ticker *time.Ticker, db *sql.DB) {
	for range ticker.C {
		err := DeleteExpiredSessions(db)
		if err != nil {
			log.Printf("Error deleting expired sessions: %v", err)
		} else {
			fmt.Println("Expired sessions deleted successfully.")
		}
	}
}

func DeleteExpiredSessions(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM sessions WHERE expiration < CURRENT_TIMESTAMP")
	return err
}
