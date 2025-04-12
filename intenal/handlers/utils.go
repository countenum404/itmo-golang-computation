package handlers

import (
	"encoding/json"
	"net/http"
)

const JSON = "application/json"

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", JSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ReadJson(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
