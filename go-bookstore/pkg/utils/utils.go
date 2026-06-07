package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// Something to save my time (SOOOO MUCH BOILERPLATE IF I HADNT WRITTEN THIS)
func ParseBody(r *http.Request, X interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, X)
}

func WriteJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func SendSuccess(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"data": data,
	})
}

// Thank me later
