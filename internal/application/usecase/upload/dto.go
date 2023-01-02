package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/lucassimon/golang-upload-api/pkg/entity"
)

type MediaOutput struct {
	Id          entity.ID
	Title       string
	Description string
	Alt         string
	Link        string
}

type MediaInput struct {
	File        *multipart.FileHeader
	Filename    string
	UniqueName  string
	Extension   string
	ContentType string
	Title       string
	Description string
	Alt         string
	Size        int64
}

func MakeMediaInput(file *multipart.FileHeader, title, description, alt []string) MediaInput {

	return MediaInput{
		File:        file,
		Filename:    file.Filename,
		UniqueName:  GenerateUniqueName(file),
		Extension:   GetExtension(file),
		ContentType: file.Header.Get("Content-Type"),
		Size:        file.Size,
		Title:       strings.Join(title, ""),
		Description: strings.Join(description, ""),
		Alt:         strings.Join(alt, ""),
	}
}

func GenerateUniqueName(fh *multipart.FileHeader) string {
	log.Printf("headers: %+v\n", fh.Header)
	originalFileName := strings.TrimSuffix(filepath.Base(fh.Filename), filepath.Ext(fh.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-")
	return fmt.Sprintf("%s-%s%s", uuid.New().String(), filename, filepath.Ext(fh.Filename))
}

func GetExtension(fh *multipart.FileHeader) string {
	return filepath.Ext(fh.Filename)
}
