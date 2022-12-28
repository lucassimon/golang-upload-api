package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type MediaOutput struct {
	Name        string
	ContentType string
	Link        string
	Size        int64
}

type MediaInput struct {
	File        *multipart.FileHeader
	Filename    string
	UniqueName  string
	Extension   string
	ContentType string
	Size        int64
}

func MakeMediaInput(file *multipart.FileHeader) MediaInput {

	return MediaInput{
		File:        file,
		Filename:    file.Filename,
		UniqueName:  generateUniqueName(file),
		Extension:   getExtension(file),
		ContentType: file.Header.Get("Content-Type"),
		Size:        file.Size,
	}
}

func generateUniqueName(fh *multipart.FileHeader) string {
	log.Printf("headers: %+v\n", fh.Header)
	originalFileName := strings.TrimSuffix(filepath.Base(fh.Filename), filepath.Ext(fh.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-")
	return fmt.Sprintf("%s-%s%s", uuid.New().String(), filename, filepath.Ext(fh.Filename))
}

func getExtension(fh *multipart.FileHeader) string {
	return filepath.Ext(fh.Filename)
}
