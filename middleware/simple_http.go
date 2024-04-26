package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func usersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// log.Println("GET")
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	if r.Method == "POST" {
		// log.Println("POST")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		users = append(users, u)

		fmt.Fprintf(w, "Hello %s created user", "POST")

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		next.ServeHTTP(w, r) // call next
// 		log.Printf("Server http middleware: %s %s %s", r.Method, r.URL, time.Since(start))
// 	}

// }

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r) // call next
	log.Printf("Server http Middleware: %s %s %s", r.Method, r.URL, time.Since(start))
}
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != "admin" || p != "password" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("username or password is incorrect!"))
			return
		}
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`Can parse the basic auth`))
			return
		}
		fmt.Println("Auth: Passed")
		next.ServeHTTP(w, r)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", authMiddleware(usersHandler))
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":80",
		Handler: logMux,
	}
	log.Println("Server Start at port name: 80")
	log.Fatal(srv.ListenAndServe())
	log.Println("Server Stop!")
}
