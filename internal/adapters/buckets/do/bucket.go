package do

import (
	"context"
	"log"
	"mime/multipart"
)

type ProviderDigitalOcean struct {
}

func NewDigitalOceanBucket(ctx context.Context) *ProviderDigitalOcean {
	return &ProviderDigitalOcean{}
}

func (b *ProviderDigitalOcean) Upload(ctx context.Context, file *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Println("uploading the file in bucket in bucket DigitalOcean")

	return "media dto or media output or media entity", nil

}
