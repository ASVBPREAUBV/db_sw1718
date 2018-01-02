package main

import (
	"net/http"
	"fmt"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Println")
	fmt.Fprintf(w, "Hello Fprintf");
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about_handler")
	fmt.Fprintf(w, "about_handler");
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	fmt.Println("Starting Server...")
	http.ListenAndServe(":8000", nil)
}

