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
	var medias []*uc.MediaOutput

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

	form := r.MultipartForm
	title := form.Value["title"]
	description := form.Value["description"]
	alt := form.Value["alt"]

	log.Println(title, description, alt)

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]

	// TODO: Verificar o uso correto do context
	ctx := context.Background()

	usecase, err := uc.NewUploadFilesUseCase(ctx, h.mediaDB)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Println(file.Filename)

		input_params := uc.MakeMediaInput(file, title, description, alt)

		media, err := usecase.Execute(ctx, input_params)

		if err != nil {
			log.Fatal(err)
		}
		medias = append(medias, media)
		log.Println("continue to next file")
	}

	msg := struct {
		Message string            `json:"message"`
		Data    []*uc.MediaOutput `json:"data"`
	}{
		Message: "ok",
		Data:    medias,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
