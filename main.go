package main

import (
	"log"
	"net/http"

	"nasa-api/handlers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("There is some error with environment variables")
	}

	http.HandleFunc("GET /apod", handlers.ApodHandler)

	log.Println("Server running on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
