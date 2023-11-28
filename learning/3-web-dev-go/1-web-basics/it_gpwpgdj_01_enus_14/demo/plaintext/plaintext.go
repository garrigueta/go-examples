package plaintext

import (
    "net/http"
	"fmt"
	"demo.com/demo"
)

type CarHandler struct {
	Car demo.Car
}
func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", h.Car.Make, h.Car.Model, h.Car.Year)
	}
	
type PersonHandler struct {
	Person demo.Person
}

func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Name: %v %v", h.Person.FirstName, h.Person.LastName)
}
	