package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateTableWhenNotExists(db *sql.DB) {
	createTb := `
	CREATE TABLE IF NOT EXISTS expense (id SERIAL  PRIMARY KEY , title TEXT , amount FLOAT , note TEXT ,tags TEXT[]);
	`
	_, err := db.Exec(createTb)

	if err != nil {
		fmt.Printf("cant create table")
	}

}

func ConnectDB() *sql.DB {

	url := os.Getenv("DATABASE_URL")
	fmt.Printf("Connecting to: " + url + "\n")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	// CreateTableWhenNotExists(db)

	return db
}
