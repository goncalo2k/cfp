package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bonjour monde")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
