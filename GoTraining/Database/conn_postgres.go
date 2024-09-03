package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	fmt.Println(db, err)
	defer db.Close()
	result, err := db.Exec("insert into dept values (2,'Fin', 'Hyd')")
	fmt.Println(result, err)
	rows, err := result.RowsAffected()
	fmt.Println(rows)
}
