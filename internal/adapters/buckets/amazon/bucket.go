package amazon

import (
	"context"
	"log"
	"mime/multipart"
)

type ProviderAmazon struct {
}

func NewAmazonBucket(ctx context.Context) (*ProviderAmazon, error) {
	return &ProviderAmazon{}, nil
}

func (b *ProviderAmazon) Upload(ctx context.Context, file *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Println("uploading the file in bucket Amazon")
	return "media dto or media output or media entity", nil

}
