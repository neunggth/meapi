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

	row := db.QueryRow("INSERT INTO users(name, age) VALUES($1, $2) RETURNING *", "Neung", 11)
	var id int
	var name string
	var age int
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("Insert data error", err)
	}
	fmt.Println(id, name, age)
}
