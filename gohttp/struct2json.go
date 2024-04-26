package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID:   1,
		Name: "Nattaphon",
		Age:  21,
	}
	b, err := json.Marshal(u)
	fmt.Printf("byte : %T \n ", b)
	fmt.Printf("err : %s \n ", b)
	fmt.Println(err)

}
