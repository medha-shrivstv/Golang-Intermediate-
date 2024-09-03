package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (my *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>First Handler</h1>")
}
func main() {

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<h1>Hello World</h1>")
			fmt.Fprintf(w, "Method %s", r.Method)
		})
	http.Handle("/first", new(MyHandler))
	fmt.Println("Starting server on 8080 ")
	http.ListenAndServe(":8080", nil)
}
