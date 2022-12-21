package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/pawutj/assessment/pkg/entities"

	"github.com/stretchr/testify/assert"
)

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	// req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

func TestGetUserByID(t *testing.T) {

	var e entities.Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(1)), nil)
	// res := request(http.MethodGet, uri("expenses", "1"), nil)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, e.Title, "Title")

}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
