package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"forum/database"
	"forum/handlers"

	// "forum/websocket"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// handlers.Info = make(map[int]handlers.User)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <database_file>")
		os.Exit(1)
	}

	// dispatcher := handlers.NewDispatcher(5) // 5 workers
    // dispatcher.Run()
    
   
    
    // Example of adding a job
    // job := handlers.Job{
    //     ID:      1,
    //     Type:    "notification",
    //     Payload: map[string]interface{}{"userId": 123},
    // }
    // dispatcher.JobQueue <- job

	fmt.Printf("Creating database at %v...\n", os.Args[1:])
	db := database.OpenDatabase(os.Args[1])
	database.CreateTables(db)
	database.CreateTriggers(db)
	fmt.Println("Database setup complete!")
	ticker := time.NewTicker(time.Hour)
	go database.DES_Ticker(ticker, db)
	defer ticker.Stop()
	defer db.Close()

	handlers.DB = db
	fmt.Println("User created successfully!")
	fmt.Println(database.QuerryPostsbyUser(db, "test", 0, 5, 0))

	http.HandleFunc("/ws", handlers.HandleWebSocket)
	http.HandleFunc("/api/chat/history", handlers.GetChatHistory)

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/infinite-scroll", handlers.InfiniteScroll)

	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("POST /register", handlers.Register)
	http.HandleFunc("/login", handlers.RegisterPage)
	http.HandleFunc("/register", handlers.RegisterPage)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/error", handlers.Error)

	http.HandleFunc("/post/{id}", handlers.GetPost)

	http.HandleFunc("/CreateComment", handlers.AddCommentHandler)
	http.HandleFunc("/createPost", handlers.CreatePost)
	http.HandleFunc("/PostReaction", handlers.PostReaction)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets"))))

	fmt.Println("Server listening on :9090...")
	fmt.Println("http://localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err.Error())
	}
}
