package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", response)
	log.Println("Serving on localhost:8123")
	http.ListenAndServe(":8123", nil)
}

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
