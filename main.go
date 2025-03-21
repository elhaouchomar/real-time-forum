package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"forum/database"
	"forum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, ticker := creatingDatabase()
	go database.DES_Ticker(ticker, db)
	defer ticker.Stop()
	defer db.Close()
	defer db.Close()
	handlers.DB = db

	MainServer := http.NewServeMux()
	MainServerHandler(MainServer)
	StartServers(MainServer)
}


func StartServers(MainServer *http.ServeMux) {
	fmt.Println("Trying runing server...")
	fmt.Println("Main Server Started:\n http://localhost:9090")
	err := http.ListenAndServe(":9090", MainServer)
	if err != nil {
		panic(err.Error())
	}
}

func MainServerHandler(Main *http.ServeMux) {

	Main.HandleFunc("/", handlers.HomePage)

	Main.HandleFunc("/ws", handlers.HandleWebSocket)
	Main.HandleFunc("/api/chat/history", handlers.GetChatHistory)
	Main.HandleFunc("POST /api/mark-read", handlers.MarkMessagesAsRead)

	Main.HandleFunc("/infinite-scroll", handlers.InfiniteScroll)

	Main.HandleFunc("POST /checker", handlers.Checker)
	Main.HandleFunc("POST /login", handlers.Login)
	Main.HandleFunc("POST /register", handlers.Register)
	Main.HandleFunc("/logout", handlers.Logout)

	Main.HandleFunc("/post/{id}", handlers.GetPost)

	Main.HandleFunc("POST /CreateComment", handlers.AddCommentHandler)
	Main.HandleFunc("POST /createPost", handlers.CreatePost)
	Main.HandleFunc("POST /PostReaction", handlers.PostReaction)

	Main.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets"))))
}

func creatingDatabase() (*sql.DB, *time.Ticker) {
	db, err := database.OpenDatabase("base.db")
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	database.CreateTables(db)
	database.CreateTriggers(db)
	fmt.Println("Database setup complete!")

	ticker := time.NewTicker(time.Hour)
	return db, ticker
}
