package main

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
	"log"
)

func createDatabase() {
    db,err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(`create database nissero_albarracin_melo_banchero_db1`)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	createDatabase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=nissero_albarracin_melo_banchero_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
