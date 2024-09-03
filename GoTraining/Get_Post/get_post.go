package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Dept represents the structure of a department
type Dept struct {
	Deptno int    `json:"deptno"`
	Dname  string `json:"dname"`
	Loc    string `json:"loc"`
}

// create a slice of Dept to hold department data
var departments []Dept

// ServeHTTP handles GET and POST requests
func (my *Dept) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		// Marshal the slice of departments to JSON
		data, err := json.Marshal(departments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Write the JSON data to the response
		w.WriteHeader(http.StatusOK)
		w.Write(data)

	case "POST":
		// Create a new Dept instance
		var newDept Dept

		// Decode the JSON request body into the newDept
		err := json.NewDecoder(r.Body).Decode(&newDept)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new department to the slice
		departments = append(departments, newDept)

		// Log the current departments slice
		fmt.Println("Updated departments slice:", departments)

		// Send a success response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newDept)
	}
}

func main() {
	// Initialize some departments
	departments = []Dept{
		{Deptno: 10, Dname: "Accounting", Loc: "New York"},
		{Deptno: 20, Dname: "Research", Loc: "Dallas"},
	}

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World</h1>")
	})

	// Dept handler
	http.Handle("/dept", new(Dept))

	// Start the server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
