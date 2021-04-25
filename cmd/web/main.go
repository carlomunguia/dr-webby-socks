package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting with server on Port 8080")

	_ = http.ListenAndServe(":8000", mux)
}
