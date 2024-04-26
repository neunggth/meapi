package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()
	createTb := `CREATE TABLE IF NOT EXISTS users
	(
		id SERIAL PRIMARY KEY,
		name TEXT ,
		age INT 

	)`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err)
	}

	fmt.Println("create table success")

}
