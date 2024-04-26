package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
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
	})
	log.Println("Server Start at port name: 80")
	log.Fatal(http.ListenAndServe(":80", nil))
	log.Println("Server Stop")
}
