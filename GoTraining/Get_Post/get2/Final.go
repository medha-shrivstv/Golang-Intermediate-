package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var empNumber int

	fmt.Println("Enter employee number (1-12 valid, 13 - 404):")
	_, err := fmt.Scan(&empNumber)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Validate empNumber
	if empNumber < 1 || empNumber > 13 {
		fmt.Println("Invalid employee number.")
		return
	}

	// If empNumber is 13, simulate a 404 error
	if empNumber == 13 {
		fmt.Println("Error: Employee not found (404).")
		return
	}

	// Make a request to get user data
	user, err := getUser(empNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print user email
	fmt.Println("User Email:", user.Email)
}

func getUser(empNumber int) (*User, error) {
	url := "https://reqres.in/api/users/" + strconv.Itoa(empNumber)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user data")
	}

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse.Data, nil
}
