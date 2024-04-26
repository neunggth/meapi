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
	stmt, err := db.Prepare("UPDATE users SET name = $1 WHERE id = $2")
	if err != nil {
		log.Fatal("Can'tPrepare statment", err)
	}
	_, err = stmt.Exec("NattaphonLove", 1)
	if err != nil {
		log.Fatal("Exec error update", err)
	}
	fmt.Println("Update success")

}
