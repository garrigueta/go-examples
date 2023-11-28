package main

import (
    "net/http"
	"fmt"
)

type PersonHandler struct{
	FirstName string
	LastName string
}

type CarHandler struct{
	Make string
	Model string
	Year int
}
func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Name: %v %v", h.FirstName, h.LastName)
	}
	
func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", h.Make, h.Model, h.Year)
	}

func main() {

	http.Handle("/person",&PersonHandler{"Bob", "Smith"})
	http.Handle("/car",&CarHandler{"Ford", "Fiesta", 2005})
    	
	http.ListenAndServe(":8081", nil)

}