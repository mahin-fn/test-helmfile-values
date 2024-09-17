package test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func TestDatabaseConnection(t *testing.T) {
	// Replace the connection parameters with your database details
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", "admin", "secret", "127.0.0.1:3306", "testdb")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	t.Log("Database connection successful")
}
