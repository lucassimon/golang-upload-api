package cloudinary

import (
	"context"
	"log"
)

type ProviderCloudinary struct {
}

func NewCloudinaryBucket(ctx context.Context) *ProviderCloudinary {
	return &ProviderCloudinary{}
}

func (b *ProviderCloudinary) Upload(ctx context.Context) (string, error) {
	log.Println("uploading the file in bucket Cloudninary")

	return "media dto or media output or media entity", nil

}
