// Package connection provides functions for connecting to a MySQL database.
package connection

import (
	"fmt"
	"log"
	"time"

	"github.com/mokhlesurr031/sre-project/backend/internal/envs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		envs.DB().Username, envs.DB().Password, envs.DB().Host, envs.DB().Port, envs.DB().Name)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	db = d // Assigning the value to the package-level db variable

	// Set up connection pool or Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// Configure connection pool settings
	sqlDB.SetMaxIdleConns(10)                  // Maximum idle connections in pool
	sqlDB.SetMaxOpenConns(100)                 // Maximum open connections
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // Maximum lifetime of a connection

	// Ping to verify the database connection
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}

// ConnectDB sets the db client of database using default configuration file
func ConnectDB() error {
	if err := initDB(); err != nil {
		return err
	}
	return nil
}

// DefaultDB returns default db
func DefaultDB() *gorm.DB {
	return db.Debug()
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	log.Println("Database Connection Closed")
	return sqlDB.Close()
}
