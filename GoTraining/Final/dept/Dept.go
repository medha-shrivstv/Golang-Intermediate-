package deptdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlePostRequest handles the incoming POST request
func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	var depts []Dept

	// Decode the JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&depts); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Process the depts slice (e.g., insert into database)
	// ...

	w.WriteHeader(http.StatusOK)
}

// GetRecords retrieves all records from the "dept" table
func GetRecords() ([]Dept, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT deptno, dname, loc FROM dept")
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	var depts []Dept
	for rows.Next() {
		var dept Dept
		if err := rows.Scan(&dept.Deptno, &dept.Dname, &dept.Loc); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		depts = append(depts, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed: %v", err)
	}

	return depts, nil
}

// InsertRecord inserts a new record into the "dept" table
func InsertRecord(dept Dept) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO dept (deptno, dname, loc) VALUES ($1, $2, $3)"
	_, err = db.Exec(query, dept.Deptno, dept.Dname, dept.Loc)
	if err != nil {
		return fmt.Errorf("failed to insert record: %v", err)
	}

	return nil
}

// UpdateRecord updates an existing record in the "dept" table
func UpdateRecord(dept Dept) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE dept SET dname = $2, loc = $3 WHERE deptno = $1"
	_, err = db.Exec(query, dept.Deptno, dept.Dname, dept.Loc)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	return nil
}

// DeleteRecord deletes a record from the "dept" table
func DeleteRecord(deptno int) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM dept WHERE deptno = $1"
	_, err = db.Exec(query, deptno)
	if err != nil {
		return fmt.Errorf("failed to delete record: %v", err)
	}

	return nil
}
