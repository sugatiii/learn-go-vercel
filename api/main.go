package handler

import (
	"encoding/json"
	"net/http"
)

// Exported function untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from Golang API on Vercel!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
