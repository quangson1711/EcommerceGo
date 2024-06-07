package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Validate = validator.New()

func ParseJson(request *http.Request, payload any) error {
	if request.Body == nil {
		return fmt.Errorf("missing body")
	}

	return json.NewDecoder(request.Body).Decode(payload)
}

func WirteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WirteError(w http.ResponseWriter, status int, err error) {
	WirteJson(w, status, map[string]string{"error": err.Error()})
}
