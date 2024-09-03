package main

import (
	"fmt"
)

// Define the Dept struct
type Depts struct {
	deptno int
	dname  string
	loc    string
}

func main() {
	// Create a slice to hold the departments
	var departments []Depts

	for {
		// Display the menu options
		fmt.Println("Menu:")
		fmt.Println("1. Insert Department")
		fmt.Println("2. List Departments")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1, 2, or 3): ")

		// Read user choice
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Insert department
			var depts Depts
			fmt.Print("Enter department number: ")
			fmt.Scan(&depts.deptno)
			fmt.Print("Enter department name: ")
			fmt.Scan(&depts.dname)
			fmt.Print("Enter location: ")
			fmt.Scan(&depts.loc)

			// Add the department to the slice
			departments = append(departments, depts)
			fmt.Println("Department added successfully!\n")

		case 2:
			// List all departments
			if len(departments) == 0 {
				fmt.Println("No departments found.\n")
			} else {
				fmt.Println("Department List:")
				for _, dept := range departments {
					fmt.Printf("Dept No: %d, Dept Name: %s, Location: %s\n", dept.deptno, dept.dname, dept.loc)
				}
				fmt.Println()
			}

		case 3:
			// Exit the program
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.\n")
		}
	}
}
