package utils

import (
	"encoding/json"
	"net/http"

	"github.com/2marks/go-expense-tracker-api/errors"
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

	_code := statusCode
	errHttp, ok := err.(*errors.ErrHttpRequest)
	if ok {
		_code = errHttp.StatusCode
	}

	return writeResponseToJson(w, _code, jsonResponse)
}
