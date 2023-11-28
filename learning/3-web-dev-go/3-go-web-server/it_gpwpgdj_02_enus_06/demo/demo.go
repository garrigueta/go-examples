package main

import (
    "net/http"
	"fmt"
)

type TestHandler struct{}
func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "This is a custom handler called TestHandler.")
	}
	
func TestHandlerFunc(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "This is a customer handler function called TestHandlerFunc.")
	}

func main() {

	http.Handle("/handler",&TestHandler{})
	http.HandleFunc("/handlerfunc",TestHandlerFunc)
    http.ListenAndServe(":8081", nil)

}