package utils

import (
	"encoding/json"
	"net/http"
)

func Res(w http.ResponseWriter, result interface{}, statusCode int) {
	userJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// Write JSON response
	_, err = w.Write(userJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
