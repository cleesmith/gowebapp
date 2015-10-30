package database

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var (
	// Database wrapper
	DB *bolt.DB
)

// Database is the details for the database connection
type Database struct {
	Name string
}

func Connect(d Database) {
	var err error
	// DB, err := bolt.Open("bolt.db", 0644, nil)
	if DB, err = bolt.Open(d.Name, 0644, nil); err != nil {
		log.Fatal(err)
	}
}
