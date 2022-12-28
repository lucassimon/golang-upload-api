package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucassimon/golang-upload-api/internal/adapters/db"
	uc "github.com/lucassimon/golang-upload-api/internal/application/usecase/upload"
)

type UploadV1Handler struct {
	mediaDB db.MediaDBRepositoryInterface
}

func NewUploadV1Handler(mediaDB db.MediaDBRepositoryInterface) *UploadV1Handler {
	return &UploadV1Handler{
		mediaDB: mediaDB,
	}
}

func (h *UploadV1Handler) Upload(w http.ResponseWriter, r *http.Request) {

	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	// form := r.MultipartForm()
	// files := form.File["file"]

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]

	for _, file := range files {
		log.Println(file.Filename)
		// TODO: convert file.Header to file
		input_params := uc.MakeMediaInput(file)

		// TODO: Verificar o uso correto do context
		ctx := context.Background()

		usecase := uc.NewUploadFilesUseCase(ctx, h.mediaDB)
		// TODO: pass input_params as arguments
		_, err := usecase.Execute(ctx, input_params)

		if err != nil {
			log.Fatal(err)
		}
		log.Println("continue to next file")
	}

	msg := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
