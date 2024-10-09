package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)

	http.ListenAndServe(":8080", nil)
}
