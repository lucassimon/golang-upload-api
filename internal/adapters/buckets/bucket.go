package buckets

import (
	"context"
	"log"

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
	Upload(context.Context) (string, error)
}

type BucketFactory struct {
	Name       string
	BucketName string
	Path       string
	Provider   ProviderInterface
}

func MakeBucketStrategy(ctx context.Context, provider string) *BucketFactory {

	switch provider {
	case "amazon":
		provider := amazon.NewAmazonBucket(ctx)
		bucket := &BucketFactory{
			Name:       "amazon",
			BucketName: "some-bucket-name/",
			Path:       "uploads/",
			Provider:   provider,
		}
		log.Println("Strategy Amazon S3 Bucket")
		return bucket
	case "do":
		provider := do.NewDigitalOceanBucket(ctx)
		bucket := &BucketFactory{
			Name:       "digitalOcean",
			BucketName: "megustaviajar",
			Path:       "uploads/",
			Provider:   provider,
		}
		log.Println("Strategy Digital Ocean Spaces")
		return bucket
	case "cloudinary":
		provider := cloudinary.NewCloudinaryBucket(ctx)
		bucket := &BucketFactory{
			Name:       "cloudinary",
			BucketName: "some-bucket-name/",
			Path:       "uploads/",
			Provider:   provider,
		}
		log.Println("Strategy Google cloud computing")
		return bucket
	case "gcp":
		provider := gcp.NewGCPBucket(ctx)
		bucket := &BucketFactory{
			Name:       "google",
			BucketName: gcp.GetBucketName(),
			Path:       "uploads/",
			Provider:   provider,
		}
		log.Println("Strategy Google cloud computing")
		return bucket

	default:
		provider := local.NewLocalBucket(ctx)
		bucket := &BucketFactory{
			Name:       "local",
			BucketName: "medias/",
			Path:       "uploads/",
			Provider:   provider,
		}
		log.Println("Strategy Local storage")
		return bucket
	}
}
