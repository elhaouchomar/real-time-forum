package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)


// In modd.conf
// **/*.go {
// perp: go test @dirmods
// }

// **/*.go !**/*_test.go {
// prep: go build -o main .
// daemon +sigterm: ./repo
// }
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "multipart/form-..")
    w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `<h1 style="color: red;">Welcome to my awsoeme site!</h1>
    <input type="text" name="test">
    <button>Send</button>`)
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `<h1>Contact page</h1><p>To get touch, email me at
    <a href="www.gmail.com">Gmail</a>.</p>`)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
		return
	case "/contact":
		contacthandler(w, r)
		return
	default:
		// TODO:  handle page not found err
	}
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// 	return
	// } else if r.URL.Path == "/contact" {
	// 	contacthandler(w, r)
	// 	return
	// }
	// handle te page not found err
}

func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/", pathHandler)
	// mux.HandleFunc("/contact", contacthandler)
	fmt.Println("Starting the server on :3000...")
    time.Sleep(1 * time.Second)
	fmt.Fprintln(os.Stdout, "http:///localhost:3000")
	err := http.ListenAndServe(":3000", mux); panic(err)
}
