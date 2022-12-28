package services

import (
	"context"
	"log"

	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets"
	"github.com/lucassimon/golang-upload-api/internal/adapters/db"
	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
)

type UploadService struct {
	Bucket     *buckets.BucketFactory
	Repository db.MediaDBRepositoryInterface
}

func NewUploadService(bucket *buckets.BucketFactory, mediaDB db.MediaDBRepositoryInterface) *UploadService {
	return &UploadService{Bucket: bucket, Repository: mediaDB}
}

func (s *UploadService) Upload(ctx context.Context) (string, error) {
	log.Println("call upload service")
	result, err := s.Bucket.Provider.Upload(ctx)

	if err != nil {
		return "", err
	}
	log.Println("media uploaded")

	return result, nil
}

func (s *UploadService) Save(ctx context.Context, media *entity.MidiaEntity) error {
	log.Println("call save service")
	err := s.Repository.Create(media)

	if err != nil {
		return err
	}
	log.Println("entity saved")
	return nil
}
