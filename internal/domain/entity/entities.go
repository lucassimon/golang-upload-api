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
	Title       string    `valid:"required" json:"title"`
	Description string    `valid:"required" json:"description"`
	Alt         string    `valid:"required" json:"alt"`
}

func MakeMediaEntity(
	name,
	contenttype,
	link,
	provider,
	bucketName,
	directory string,
	size int64,
	title, description, alt string,
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
		Title:       title,
		Description: description,
		Alt:         alt,
	}
}

func (m *MediaEntity) Print() string {
	return fmt.Sprintf("{Id:%s, Title:%s, Link:%s}", m.Id, m.Title, m.Link)
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

func (m *MediaEntity) GetTitle() string {
	return m.Title
}

func (m *MediaEntity) GetDescription() string {
	return m.Description
}

func (m *MediaEntity) GetAlt() string {
	return m.Alt
}
