package local

import (
	"context"
	"log"
)

type ProviderLocal struct {
}

func NewLocalBucket(ctx context.Context) *ProviderLocal {
	return &ProviderLocal{}
}

func (b *ProviderLocal) Upload(ctx context.Context) (string, error) {
	log.Println("uploading the file in bucket Local")
	return "media dto or media output or media entity", nil

}
