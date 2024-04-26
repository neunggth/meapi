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

	stmt, err := db.Prepare("SELECT * FROM users")
	if err != nil {
		log.Fatal("Prepare error", err)

	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query error", err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Can't Scan row into variable", err)
		}
		fmt.Println(id, name, age)

	}
	defer stmt.Close()
}
