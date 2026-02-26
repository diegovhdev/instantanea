package helpers

import (
	"encoding/json"
	"net/http"
)


func WriteJSON(w http.ResponseWriter, statusCode int, s any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(s)
}