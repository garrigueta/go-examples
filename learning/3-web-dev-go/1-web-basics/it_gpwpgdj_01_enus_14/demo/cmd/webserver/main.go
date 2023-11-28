package main

import (
    "net/http"
	"demo.com/demo"
	"demo.com/demo/html"
	"demo.com/demo/plaintext"
)

func main() {
	
	http.Handle("/person/html",&html.PersonHandler{demo.Person{"Bob", "Smith"}})
	http.Handle("/car/html",&html.CarHandler{demo.Car{"Ford", "Fiesta", 2005}})

	http.Handle("/person/plaintext",&plaintext.PersonHandler{demo.Person{"Bob", "Smith"}})
	http.Handle("/car/plaintext",&plaintext.CarHandler{demo.Car{"Ford", "Fiesta", 2005}})
    	
	http.ListenAndServe(":8081", nil)

}