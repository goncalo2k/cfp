package main

import (
	"fmt"
	"net/http"

	"backend/utils"

	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bonjour monde")
}

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on port:", utils.GetEnvVar("PORT"))
	http.ListenAndServe(":"+utils.GetEnvVar("PORT"), nil)
}
