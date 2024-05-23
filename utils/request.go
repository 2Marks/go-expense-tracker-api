package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseRequestBody[T any](r *http.Request, payload *T) error {
	body := r.Body

	if body == nil {
		return fmt.Errorf("no payload sent in request body")
	}

	return json.NewDecoder(body).Decode(payload)
}
