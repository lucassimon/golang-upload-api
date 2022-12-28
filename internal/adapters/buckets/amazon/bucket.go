package amazon

import (
	"context"
	"log"
)

type ProviderAmazon struct {
}

func NewAmazonBucket(ctx context.Context) *ProviderAmazon {
	return &ProviderAmazon{}
}

func (b *ProviderAmazon) Upload(ctx context.Context) (string, error) {
	log.Println("uploading the file in bucket Amazon")
	return "media dto or media output or media entity", nil

}
