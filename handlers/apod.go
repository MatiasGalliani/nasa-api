package handlers

import (
	"encoding/json"
	"net/http"

	"nasa-api/services"
)

func ApodHandler(response http.ResponseWriter, request *http.Request) {
	apod, err := services.GetApod()
	if err != nil {
		http.Error(response, "Error while obtaining NASA data", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(apod)
}
