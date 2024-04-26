package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Nattaphon", Age: 21},
	{ID: 2, Name: "Thirawat", Age: 22},
	{ID: 3, Name: "Nattaphon", Age: 23},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

// GET
func getUserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

// POST
func createUserHandler(c echo.Context) error {

	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	users = append(users, u)
	return c.JSON(http.StatusCreated, u)

}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	createTb := `CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, name TEXT, age INT)`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err)
	}
	fmt.Println("create table success")
}
func main() {
	InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", healthHandler)

	g := e.Group("/api")
	// auth
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "password" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/users", getUserHandler)
	g.GET("/users/:id", getUserHandler)
	g.POST("users", createUserHandler)

	log.Println("Server Start at port name: 80")

	log.Fatal(e.Start(":80"))
	log.Println("Server Stop!")
}
