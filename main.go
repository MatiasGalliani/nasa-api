package main

import (
	"log"
	"net/http"

	"nasa-api/handlers"
	"nasa-api/middleware"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("There is some error with environment variables")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /apod", handlers.ApodHandler)

	log.Println("Server running on :8000")
	if err := http.ListenAndServe(":8000", middleware.CORS(mux)); err != nil {
		log.Fatal(err)
	}
}
