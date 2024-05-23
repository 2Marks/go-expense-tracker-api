package utils

import (
	"encoding/json"
	"net/http"
)

func writeResponseToJson(w http.ResponseWriter, statusCode int, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(response)
}

func WriteSuccessResponseToJson(w http.ResponseWriter, statusCode int, message string, data interface{}) error {
	jsonResponse := map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}

	return writeResponseToJson(w, statusCode, jsonResponse)
}

func WriteErroresponseToJson(w http.ResponseWriter, statusCode int, err error) error {
	jsonResponse := map[string]interface{}{
		"success": false,
		"message": err.Error(),
	}

	return writeResponseToJson(w, statusCode, jsonResponse)
}
