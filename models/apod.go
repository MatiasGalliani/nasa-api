package models

type Apod struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	URL         string `json:"url"`
	Date        string `json:"date"`
}
