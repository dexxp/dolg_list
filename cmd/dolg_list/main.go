package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!123")
}

func main() {
	http.HandleFunc("/", helloHandler)

    http.ListenAndServe(":8080", nil)
}
