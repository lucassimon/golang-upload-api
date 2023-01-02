package cloudinary

import (
	"context"
	"log"
	"mime/multipart"
)

type ProviderCloudinary struct {
}

func NewCloudinaryBucket(ctx context.Context) (*ProviderCloudinary, error) {
	return &ProviderCloudinary{}, nil
}

func (b *ProviderCloudinary) Upload(ctx context.Context, file *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Println("uploading the file in bucket Cloudninary")

	return "media dto or media output or media entity", nil

}
