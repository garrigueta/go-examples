package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	testHTTPCal := func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "Tis is a test response.")
	}

	request := httptest.NewRequest("GET", "http://testurl.com", nil)
	recorder := httptest.NewRecorder()
	testHTTPCal(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(string(responseBody))
}
