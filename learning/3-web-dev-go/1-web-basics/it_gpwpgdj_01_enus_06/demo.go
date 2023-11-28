package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("https://sp-pp.travelport.com")

	responseBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(string(responseBody))
}
