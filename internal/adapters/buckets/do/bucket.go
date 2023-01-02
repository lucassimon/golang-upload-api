package do

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ProviderDigitalOcean struct {
	session    *session.Session
	client     *s3.S3
	bucketName string
	region     string
}

func NewDigitalOceanBucket(ctx context.Context) (*ProviderDigitalOcean, error) {
	spaceRegion := getDigitalOceanSpaceRegion()
	accessKey := getDigitalOceanAccessKey()
	secretKey := getDigitalOceanSecretKey()
	endpoint := getDigitalOceanSpaceEndpoint()
	bucketName := getDigitalOceanSpaceName()

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(spaceRegion),
	}

	newSession, err := session.NewSession(s3Config)

	if err != nil {
		return nil, err
	}
	s3Client := s3.New(newSession)

	return &ProviderDigitalOcean{
		session:    newSession,
		client:     s3Client,
		bucketName: bucketName,
		region:     spaceRegion,
	}, nil
}

func getDigitalOceanAccessKey() string {
	return os.Getenv("DO_ACCESS_KEY")
}

func getDigitalOceanSecretKey() string {
	return os.Getenv("DO_SECRET_KEY")
}

func getDigitalOceanSpaceName() string {
	return os.Getenv("DO_SPACE_NAME")
}

func getDigitalOceanSpaceRegion() string {
	return os.Getenv("DO_SPACE_REGION")
}

func getDigitalOceanSpaceEndpoint() string {
	return os.Getenv("DO_SPACE_ENDPOINT")
}

func (b *ProviderDigitalOcean) Upload(ctx context.Context, fh *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Println("uploading the file in bucket in bucket DigitalOcean")

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buff := make([]byte, fh.Size)
	_, err = file.Read(buff)
	if err != nil {
		return "", err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	fileBytes := bytes.NewReader(buff)

	key := fmt.Sprintf("uploads/%s", uniqueName)

	object := s3.PutObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(key),
		Body:   fileBytes,
		ACL:    aws.String("public-read"),
	}

	_, uploadErr := b.client.PutObject(&object)
	// log.Printf("created entity %+v\n", _)

	if uploadErr != nil {
		return "", fmt.Errorf("error uploading file: %v", uploadErr)
	}

	fileUrl := fmt.Sprintf("https://%v.%v.digitaloceanspaces.com/%v", b.bucketName, b.region, key)
	return fileUrl, nil
}
