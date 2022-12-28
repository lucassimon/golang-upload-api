package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassimon/golang-upload-api/internal/adapters/db"
)

type MediasHandler struct {
	mediaDB db.MediaDBRepositoryInterface
}

func NewMediasHandler(mediaDB db.MediaDBRepositoryInterface) *MediasHandler {
	return &MediasHandler{
		mediaDB: mediaDB,
	}
}

func (h *MediasHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func (h *MediasHandler) GetById(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func (h *MediasHandler) Delete(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}

func (h *MediasHandler) Sync(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
