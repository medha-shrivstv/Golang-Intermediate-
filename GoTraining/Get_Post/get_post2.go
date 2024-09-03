package main

import (
	"fmt"
	"net/http"
)

// MyHandler is a custom HTTP handler.
type MyHandler struct{}

// ServeHTTP handles HTTP requests for MyHandler.
func (my *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>First Handler</h1>")
}

// handleGetRequest handles GET requests.
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "GET request invoked")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// handlePostRequest handles POST requests.
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "POST request invoked")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Default handler for the root path.
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<h1>Hello World</h1>")
			fmt.Fprintf(w, "Method %s", r.Method)
		})

	// Custom handler for the /first path.
	http.Handle("/first", new(MyHandler))

	// Handlers for GET and POST requests.
	http.HandleFunc("/get", handleGetRequest)
	http.HandleFunc("/post", handlePostRequest)

	fmt.Println("Starting server on 8080")
	http.ListenAndServe(":8080", nil)
}
