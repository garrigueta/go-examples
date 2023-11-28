package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	testServer:= httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
		responseWriter.WriteHeader(http.StatusNotFound)
		io.WriteString(responseWriter , "This is a test response.")
	}))
	defer testServer.Close()

	response, _:= http.Get(testServer.URL)
	
	responseBody, _:= io.ReadAll(response.Body)
	
	fmt.Println(response.StatusCode)
	fmt.Println(string(responseBody))
}
