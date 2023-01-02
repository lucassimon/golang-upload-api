package upload

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

func (s *UploadService) Upload(ctx context.Context, input_params MediaInput) (string, error) {
	log.Println("call upload service")
	file := input_params.File
	uniqueName := input_params.UniqueName
	extension := input_params.Extension

	result, err := s.Bucket.Provider.Upload(ctx, file, uniqueName, extension)

	if err != nil {
		return "", err
	}
	log.Println("media uploaded")

	return result, nil
}

func (s *UploadService) Save(ctx context.Context, media *entity.MediaEntity) error {
	log.Println("call save service")
	_, err := s.Repository.Create(media)

	if err != nil {
		return err
	}
	log.Println("entity saved")
	return nil
}
