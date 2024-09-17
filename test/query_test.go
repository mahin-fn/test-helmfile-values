package test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type Resource struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func TestInsertAndFetchResource(t *testing.T) {
	// Connect to the database
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", "admin", "secret", "127.0.0.1:3306", "testdb")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Insert a record into the database
	_, err = db.Exec("INSERT INTO resources (name, url, created_at) VALUES (?, ?, ?)", "TestResource", "http://example.com", time.Now())
	if err != nil {
		t.Fatalf("Error inserting record into database: %v", err)
	}

	// Fetch the inserted record from the database
	// Fetch the inserted record from the database
	var resource Resource
	var createdAt []uint8
	err = db.QueryRow("SELECT id, name, url, created_at FROM resources WHERE name = ?", "TestResource").Scan(&resource.Id, &resource.Name, &resource.URL, &createdAt)
	if err != nil {
		t.Fatalf("Error fetching record from database: %v", err)
	}

	resource.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		t.Fatalf("Error parsing created_at value: %v", err)
	}
}
