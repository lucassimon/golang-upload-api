package buckets

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets/amazon"
	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets/cloudinary"
	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets/do"
	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets/gcp"
	"github.com/lucassimon/golang-upload-api/internal/adapters/buckets/local"
)

const (
	IMAGES_PREFIX  = "images/"
	VIDEOS_PREFIX  = "videos/"
	UPLOADS_PREFIX = "uploads/"
)

type ProviderInterface interface {
	Upload(ctx context.Context, file *multipart.FileHeader, uniqueName string, extension string) (string, error)
}

type BucketFactory struct {
	Name       string
	BucketName string
	Path       string
	Provider   ProviderInterface
}

func MakeBucketStrategy(ctx context.Context, provider string) (*BucketFactory, error) {

	switch provider {
	case "amazon":
		provider, err := amazon.NewAmazonBucket(ctx)
		if err != nil {
			return nil, err
		}
		bucket := &BucketFactory{
			Name:       "amazon",
			BucketName: "some-bucket-name/",
			Path:       UPLOADS_PREFIX,
			Provider:   provider,
		}
		log.Println("Strategy Amazon S3 Bucket")
		return bucket, nil
	case "do":
		provider, err := do.NewDigitalOceanBucket(ctx)

		if err != nil {
			return nil, err
		}

		bucket := &BucketFactory{
			Name:       "digitalOcean",
			BucketName: "megustaviajar",
			Path:       UPLOADS_PREFIX,
			Provider:   provider,
		}
		log.Println("Strategy Digital Ocean Spaces")
		return bucket, nil
	case "cloudinary":
		provider, err := cloudinary.NewCloudinaryBucket(ctx)
		if err != nil {
			return nil, err
		}
		bucket := &BucketFactory{
			Name:       "cloudinary",
			BucketName: "some-bucket-name/",
			Path:       UPLOADS_PREFIX,
			Provider:   provider,
		}
		log.Println("Strategy Google cloud computing")
		return bucket, nil
	case "gcp":
		provider, err := gcp.NewGCPBucket(ctx)
		if err != nil {
			return nil, err
		}
		bucket := &BucketFactory{
			Name:       "google",
			BucketName: gcp.GetBucketName(),
			Path:       UPLOADS_PREFIX,
			Provider:   provider,
		}
		log.Println("Strategy Google cloud computing")
		return bucket, nil

	default:
		provider, err := local.NewLocalBucket(ctx)
		if err != nil {
			return nil, err
		}
		bucket := &BucketFactory{
			Name:       "local",
			BucketName: "medias/",
			Path:       UPLOADS_PREFIX,
			Provider:   provider,
		}
		log.Println("Strategy Local storage")
		return bucket, nil
	}
}
