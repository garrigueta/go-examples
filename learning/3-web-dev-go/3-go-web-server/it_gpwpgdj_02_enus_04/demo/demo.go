package main

import (
		"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/users/billb/demo"))
	
	http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))
	
	http.HandleFunc("/samplepdf", func(res http.ResponseWriter, req *http.Request){
	http.ServeFile(res, req, "/users/billb/demo/pdf/sample.pdf")
})
	
	http.ListenAndServe(":9000", nil)
}
