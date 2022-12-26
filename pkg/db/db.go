package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	url := os.Getenv("DATABASE_URL")
	fmt.Printf("Connecting to: " + url + "\n")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
