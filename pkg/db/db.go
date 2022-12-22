package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	url := fmt.Sprintf("postgres://root:root@db:%s/go-example-db?sslmode=disable", os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
