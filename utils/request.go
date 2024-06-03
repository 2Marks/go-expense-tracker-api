package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseRequestBody[T any](r *http.Request, payload *T) error {
	body := r.Body

	if body == nil {
		return fmt.Errorf("no payload sent in request body")
	}

	return json.NewDecoder(body).Decode(payload)
}

func GetRequestQueryIntVal(r *http.Request, key string, fallback int) int {
	queryVal := r.URL.Query().Get(key)
	if queryVal == "" && fallback > 0 {
		return fallback
	}

	queryValInt, err := strconv.Atoi(queryVal)
	if err != nil {
		fmt.Printf("error in GetRequestQueryIntVal. key:%s, err:%s \n", key, err.Error())

		if fallback > 0 {
			return fallback
		}

		return 0
	}

	return queryValInt
}

func GetReqPathIntVal(r *http.Request, key string) int {
	vars := mux.Vars(r)

	val, err := strconv.Atoi(vars[key])
	if err != nil {
		fmt.Printf("error in GetReqPathIntVal. key:%s, err:%s \n", key, err.Error())
		return 0
	}

	return val
}
