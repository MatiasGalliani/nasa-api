package services

import (
	"encoding/json"
	"net/http"
	"os"

	"nasa-api/models"
)

func GetApod() (*models.Apod, error) {
	apiKey := os.Getenv("NASA_API_KEY")

	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + apiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apod models.Apod
	if err := json.NewDecoder(resp.Body).Decode(&apod); err != nil {
		return nil, err
	}

	return &apod, nil
}
