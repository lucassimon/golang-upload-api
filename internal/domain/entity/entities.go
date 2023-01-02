package entity

import (
	"fmt"

	"github.com/lucassimon/golang-upload-api/pkg/entity"
)

// sqlite>create table IF NOT EXISTS medias(id text primary key, name text, content_type text, link text, provider text, bucket_name text, directory text, size integer);

type MediaEntity struct {
	Id          entity.ID `valid:"uuidv4" json:"id"`
	Name        string    `valid:"required" json:"name"`
	ContentType string    `valid:"required" json:"content_type"`
	Link        string    `valid:"required" json:"link"`
	Provider    string    `valid:"required" json:"provider"`
	BucketName  string    `valid:"required" json:"bucket_name"`
	Directory   string    `valid:"required" json:"directory"`
	Size        int64     `valid:"required" json:"size"`
}

func MakeMediaEntity(
	name,
	contenttype,
	link,
	provider,
	bucketName,
	directory string,
	size int64,
) *MediaEntity {
	return &MediaEntity{
		Id:          entity.NewID(),
		Name:        name,
		ContentType: contenttype,
		Link:        link,
		Size:        size,
		Provider:    provider,
		BucketName:  bucketName,
		Directory:   directory,
	}
}

func (m *MediaEntity) Print() string {
	return fmt.Sprintf("{Id:%d, Title:%s, Name:%s}", m.Id, m.Name, m.ContentType)
}

func (m *MediaEntity) GetID() string {
	return m.Id.String()
}

func (m *MediaEntity) GetName() string {
	return m.Name
}

func (m *MediaEntity) GetContentType() string {
	return m.ContentType
}

func (m *MediaEntity) GetLink() string {
	return m.Link
}

func (m *MediaEntity) GetBucketName() string {
	return m.BucketName
}

func (m *MediaEntity) GetProvider() string {
	return m.Provider
}

func (m *MediaEntity) GetDirectory() string {
	return m.Directory
}

func (m *MediaEntity) GetSize() int64 {
	return m.Size
}
