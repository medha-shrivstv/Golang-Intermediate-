package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Defer a function to recover from any panic
	defer handlePanic()

	// Ask user for the page number
	var pageNumber int
	fmt.Print("Enter page number (1 or 2): ")
	fmt.Scan(&pageNumber)

	if pageNumber < 1 || pageNumber > 2 {
		panic("Invalid page number. Please enter 1 or 2.")
	}

	// Fetch data from API
	url := "https://reqres.in/api/users?page=" + strconv.Itoa(pageNumber)
	response, err := http.Get(url)
	if err != nil {
		panic("Failed to fetch data from API")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Received non-200 response code from API")
	}

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Failed to read API response")
	}

	// Parse the JSON response
	var userPage UserPage
	err = json.Unmarshal(body, &userPage)
	if err != nil {
		panic("Failed to parse API response JSON")
	}

	// Save data to a file
	err = saveToFile("demo.txt", userPage)
	if err != nil {
		panic("Failed to save data to file")
	}

	fmt.Println("Data saved successfully to demo.txt")
}

// saveToFile saves the user data to a file
func saveToFile(filename string, data UserPage) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert data to JSON format for saving
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// handlePanic is a common recovery function to handle any panics
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}
