package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)



type Todo struct {
	Title string
	Done bool
}

type TodoPage struct {
	PageTitle string
	Todos []Todo
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to index page")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPage {
		PageTitle: "All Todos",
		Todos: []Todo{
			{Title: "Learn Golang", Done: true},
			{Title: "Learn HTML", Done: true},
			{Title: "Learn CSS", Done: true},
			{Title: "Learn JavaScript", Done: true},
			{Title: "Learn Rust", Done: false},
			{Title: "Learn Python", Done: false},
			{Title: "Learn Docker", Done: true},
			{Title: "Learn Sql", Done: true},
			{Title: "Learn React.js", Done: false},
			{Title: "Learn Node.js", Done: false},
			{Title: "Learn Next.js", Done: false},
			{Title: "Learn Git", Done: true},
		},
	}
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/todos", todosHandler)
	fmt.Println("http://localhost:5050")
	http.ListenAndServe(":5050", nil)
}