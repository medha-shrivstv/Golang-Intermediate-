package deptdb

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// readProperties reads the properties from a file and returns them as a map
func readProperties(filepath string) (map[string]string, error) {
	properties := make(map[string]string)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open properties file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments and empty lines
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // ignore malformed lines
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		properties[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading properties file: %v", err)
	}

	return properties, nil
}

// getCon establishes a connection to the PostgreSQL database using properties from a file.
// It returns a pointer to sql.DB and an error if the connection fails.
func getCon(url, login, pass, dbname string) (*sql.DB, error) {
	// Add sslmode=disable to the DSN to disable SSL
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", login, pass, url, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("problem opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to establish a database connection: %v", err)
	}

	return db, nil
}

// getDb loads the database connection properties from the "dbconn.properties" file
// and establishes a connection to the database. It returns a pointer to sql.DB and an error if the connection fails.
// getDb loads the database connection properties from the "dbconn.properties" file
// and establishes a connection to the database. It returns a pointer to sql.DB and an error if the connection fails.
func getDb() (*sql.DB, error) {
	propertiesPath := "dbconn.properties"

	p, err := readProperties(propertiesPath)
	if err != nil {
		return nil, err
	}

	url, ok := p["url"]
	if !ok {
		return nil, fmt.Errorf("missing 'url' in properties file")
	}

	login, ok := p["login"]
	if !ok {
		return nil, fmt.Errorf("missing 'login' in properties file")
	}

	password, ok := p["password"]
	if !ok {
		return nil, fmt.Errorf("missing 'password' in properties file")
	}

	dbname, ok := p["dbname"]
	if !ok {
		return nil, fmt.Errorf("missing 'dbname' in properties file")
	}

	return getCon(url, login, password, dbname)
}
