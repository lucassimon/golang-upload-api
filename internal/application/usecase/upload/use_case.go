package upload

import (
	"context"
	"log"
	"os"

	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets"
	"github.com/lucassimon/golang-upload-api/internal/adapters/db"
	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
)

type UploadFilesUseCase struct {
	Service *UploadService
}

func NewUploadFilesUseCase(ctx context.Context, mediaDB db.MediaDBRepositoryInterface) (*UploadFilesUseCase, error) {
	storage := os.Getenv("STORAGE_SELECTED")
	log.Println("storage selected:", storage)

	bucket, err := buckets.MakeBucketStrategy(ctx, storage)
	if err != nil {
		return nil, err
	}
	log.Println("bucket selected:", bucket.Name)

	service := NewUploadService(bucket, mediaDB)
	log.Println("created a UploadService")

	return &UploadFilesUseCase{Service: service}, nil
}

func (uc *UploadFilesUseCase) Execute(ctx context.Context, input_params MediaInput) (*MediaOutput, error) {

	mediaUploaded, err := uc.Service.Upload(ctx, input_params)

	if err != nil {
		return nil, err
	}

	media := entity.MakeMediaEntity(
		input_params.UniqueName,
		input_params.ContentType,
		mediaUploaded,
		uc.Service.Bucket.Name,
		uc.Service.Bucket.BucketName,
		uc.Service.Bucket.Path,
		input_params.Size,
		input_params.Title,
		input_params.Description,
		input_params.Alt,
	)
	log.Printf("created entity %+v\n", media)

	log.Println("save item in database except file field")
	err = uc.Service.Save(ctx, media)

	if err != nil {
		return nil, err
	}

	log.Println("return media output or error")

	media_output := &MediaOutput{
		Id:          media.Id,
		Title:       media.Title,
		Description: media.Description,
		Alt:         media.Alt,
		Link:        media.Link,
	}
	return media_output, nil
}
