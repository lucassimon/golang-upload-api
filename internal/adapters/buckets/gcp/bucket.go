package gcp

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
)

type ProviderGCP struct {
	handler *storage.BucketHandle
}

func NewGCPBucket(ctx context.Context) (*ProviderGCP, error) {
	cli, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ProviderGCP{handler: cli.Bucket(GetBucketName())}, nil
}

// func (b *ProviderGCP) setReadOnlyAccessPublic(ctx context.Context, objectName string) {
// 	acl := b.handler.Object(objectName).ACL()
// 	err := acl.Set(ctx, storage.AllUsers, storage.RoleReader)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (b *ProviderGCP) getMidias(ctx context.Context, prefix string) ([]*dto.MidiaDTO, error) {

// 	q := &storage.Query{
// 		Prefix:    prefix,
// 		Delimiter: "/",
// 	}

// 	it := b.handler.Objects(ctx, q)

// 	midias := []*dto.MidiaDTO{}

// 	for {
// 		attrs, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}

// 		if err != nil {
// 			return nil, err
// 		}

// 		b.setReadOnlyAccessPublic(ctx, attrs.Name)

// 		midias = append(midias, buildMidia(attrs))
// 	}

// 	return midias, nil
// }

// func buildMidia(attrs *storage.ObjectAttrs) *dto.MidiaDTO {
// 	m := &dto.MidiaDTO{
// 		Name: attrs.Name,
// 		Type: attrs.ContentType,
// 		Link: "#",
// 		Size: attrs.Size,
// 	}

// 	link, err := makePublicLink(m.Name)
// 	if err != nil {
// 		log.Println(err)

// 		return m
// 	}

// 	m.Link = link
// 	return m
// }

// func makePublicLink(objectName string) (string, error) {
// 	key, err := ioutil.ReadFile(getCredentials())
// 	if err != nil {
// 		return "", err
// 	}

// 	cfg, err := google.JWTConfigFromJSON(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	urlOptions := &storage.SignedURLOptions{
// 		GoogleAccessID: cfg.Email,
// 		PrivateKey:     cfg.PrivateKey,
// 		Expires:        time.Now().Add(time.Minute * 10),
// 		Method:         http.MethodGet,
// 	}

// 	return storage.SignedURL(getBucketName(), objectName, urlOptions)
// }

// func (b *ProviderGCP) GetImages(ctx context.Context) ([]*dto.MidiaDTO, error) {
// 	images, err := b.getMidias(ctx, images_prefix)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return filterMidias(images, "image"), nil
// }

// func (b *ProviderGCP) GetVideos(ctx context.Context) ([]*dto.MidiaDTO, error) {
// 	videos, err := b.getMidias(ctx, videos_prefix)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return filterMidias(videos, "video"), nil
// }

// func (b *ProviderGCP) GetUploads(ctx context.Context) ([]*dto.MidiaDTO, error) {
// 	uploads, err := b.getMidias(ctx, uploads_prefix)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filtered := filterMidias(uploads, "image")
// 	filtered = append(filtered, filterMidias(uploads, "video")...)

// 	return filtered, nil
// }

// func filterMidias(midias []*dto.MidiaDTO, contentType string) []*dto.MidiaDTO {
// 	filtered := []*dto.MidiaDTO{}
// 	for _, m := range midias {
// 		if strings.Contains(m.Type, contentType) {
// 			filtered = append(filtered, m)
// 		}
// 	}

// 	return filtered
// }

func (b *ProviderGCP) Upload(ctx context.Context, fh *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Println("uploading the file in bucket GCP")
	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := b.handler.Object(fmt.Sprintf("uploads/%s", uniqueName)).NewWriter(ctx)
	writer.ContentType = fh.Header.Get("Content-Type")
	writer.CacheControl = "public, max-age=86400"
	writer.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	if _, err := io.Copy(writer, file); err != nil {
		return "#", err
	}
	if err := writer.Close(); err != nil {
		return "#", err
	}

	const publicLink = "https://storage.googleapis.com/%s/uploads/%s"

	return fmt.Sprintf(publicLink, GetBucketName(), uniqueName), nil
}

// func getCredentials() string {
// 	return os.Getenv("GOOGLE_CREDENTIALS")
// }

func GetBucketName() string {
	return os.Getenv("GOOGLE_BUCKET_NAME")
}
