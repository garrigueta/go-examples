package main

import (
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
    })

    http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request){
		
        http.ServeFile(w, r, "html/demopath1.html")
    })
	
	http.HandleFunc("/demopath1/subpatha", func(w http.ResponseWriter, r *http.Request){
		
        http.ServeFile(w, r, "html/demopath1_subpatha.html")
    })
	
	http.HandleFunc("/demopath2/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "html/demopath2.html")
    })

    http.ListenAndServe(":8081", nil)

}