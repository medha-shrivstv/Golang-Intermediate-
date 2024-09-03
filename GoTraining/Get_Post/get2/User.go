// User.go
package main

// User represents a single user from the API
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// UserPage represents the paginated response from the API
type UserPage struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []User `json:"data"`
	Support    struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}
type ApiResponse struct {
	Data User `json:"data"`
}
