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
	data := []byte(`{"id": 1, "name": "John Snow", "age": 21}`)
	u := &User{}
	err := json.Unmarshal(data, u)
	fmt.Printf("% #v\n", u)
	fmt.Println(err)

}
