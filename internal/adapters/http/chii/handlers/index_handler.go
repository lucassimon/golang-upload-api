package handlers

import (
	"encoding/json"
	"net/http"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	// Return the bid
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
