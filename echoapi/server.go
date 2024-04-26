package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
func main() {
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
	g.POST("users", createUserHandler)

	log.Println("Server Start at port name: 80")

	log.Fatal(e.Start(":80"))
	log.Println("Server Stop!")
}
