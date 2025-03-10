package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mattn/go-sqlite3"
)

func init() {
	sql.Register("sqlite3_with_fk", driverFunc())
}

func driverFunc() driver.Driver {
	return &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			_, err := conn.Exec("PRAGMA foreign_keys = ON;", nil)
			return err
		},
	}
}

func OpenDatabase(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3_with_fk", file)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database opened successfully!")
	return db, nil
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
