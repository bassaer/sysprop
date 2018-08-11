package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter sample"))
}

func main() {
	fmt.Println("Serving...\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
