package main

import (
	"fmt"
	"net/http"
	"time"
)

func now() string {
	return time.Now().Format(time.Stamp) + " "
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler2")
}

func logger(f http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(now() + "before")
		defer fmt.Println(now() + "after")
		f(w, r)
	}
}
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler3")
}

func main() {
	http.HandleFunc("/h1",  logger(handler1))
	http.HandleFunc("/h2",  logger(handler2))
	http.HandleFunc("/h3", logger(handler3))
	http.ListenAndServe(":5050", nil)
}
