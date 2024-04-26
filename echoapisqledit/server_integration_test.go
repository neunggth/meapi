//go build integration

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	var c User
	body := bytes.NewBufferString(`{
		"name": "John Test",
		"age": 30}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't insert data", err)
	}
	var us []User
	res := request(http.MethodGet, uri("users"), nil)
	err := res.Decode(&us)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(us), 0)

}

func uri(paths ...string) string {
	host := "http://localhost:8080/"
	if paths == nil {
		return host
	}
	url := append([]string{host}, paths...)
	return string.Join(url, "/")
}

type Response struct {
	*http.Response
	err error
}

// Decode
func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	// json unmarshal !
	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}

}
