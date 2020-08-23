package utils

import (
	"encoding/json"
	"net/http"
)

// CustomError : The error format in api response
type CustomError struct {
	Error string `json:"error"`
}

// Send : General function to send api response
func Send(w http.ResponseWriter, status int, payload interface{}) {
	result, err := json.Marshal(&payload)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

// Fail : General function to send api error response
func Fail(w http.ResponseWriter, status int, details string) {
	response := &CustomError{
		Error: details,
	}
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}
