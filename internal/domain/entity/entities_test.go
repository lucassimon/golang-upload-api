package entity

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CallMakeForTests() *MediaEntity {
	name := "Foo"
	contentType := "image/png"
	link := "http://some-url.com"
	provider := "amazon"
	bucket_name := "Foo"
	directory := "/uploads/some-file.png"
	var size int64 = 8045687521
	title := "Some Title Image"
	description := "SOme description"
	alt := "Some alternative"

	media := MakeMediaEntity(
		name,
		contentType,
		link,
		provider,
		bucket_name,
		directory,
		size,
		title,
		description,
		alt,
	)

	return media
}

func TestMedia_MakeMediaEntity(t *testing.T) {
	media := CallMakeForTests()
	require.Equal(t, "Foo", media.Name)
}

func TestMedia_Print(t *testing.T) {
	media := MediaEntity{}
	media.Id = uuid.New()
	media.Name = "Foo"
	media.ContentType = "image/png"
	media.Link = "http://some-url.com"
	media.Provider = "amazon"
	media.BucketName = "Foo"
	media.Directory = "/uploads/some-file.png"
	media.Size = 8045687521
	media.Title = "Some Title Image"
	media.Description = "SOme description"
	media.Alt = "Some alternative"

	expected := fmt.Sprintf("{Id:%s, Title:%s, Link:%s}", media.Id, media.Title, media.Link)

	result := media.Print()

	require.Equal(t, result, expected)
}

func TestMedia_GetID(t *testing.T) {
	// type ID = uuid.UUID
	// media := entity.MediaEntity{}

	// pk := uuid.MustParse("dfef6a2e-ab23-4d56-ab4c-c6173ccdd12f")
	// media.Id = pk

	// require.Equal(t, media.GetID(), pk)
}
func TestMedia_GetName(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetName(), "Foo")
}
func TestMedia_GetContentType(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetContentType(), "image/png")
}
func TestMedia_GetLink(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetLink(), "http://some-url.com")
}
func TestMedia_GetBucketName(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetBucketName(), "Foo")
}
func TestMedia_GetProvider(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetProvider(), "amazon")
}
func TestMedia_GetDirectory(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetDirectory(), "/uploads/some-file.png")
}
func TestMedia_GetSize(t *testing.T) {
	media := CallMakeForTests()

	require.Equal(t, media.GetSize(), int64(8045687521))
}
