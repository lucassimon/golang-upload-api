package db

import (
	"database/sql"
	"log"
	"testing"

	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createMedia(Db)
}

func createTable(db *sql.DB) {
	table := `
	CREATE TABLE IF NOT EXISTS medias(
		id text PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		content_type text NOT NULL,
		link text NOT NULL,
		provider text NOT NULL,
		bucket_name text NOT NULL,
		directory text NOT NULL,
		size integer NOT NULL,
		title text NOT NULL,
		description text NOT NULL,
		alt text NOT NULL
	);
	`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createMedia(db *sql.DB) {
	insert := `
	INSERT INTO medias VALUES(
		"4848560e-b438-4076-8613-4f263d743bf8",
		"0ef22762-de0c-4038-9158-8871fb431aca-16.1-background.png",
		"image/png",
		"https://megustaviajar.sfo2.digitaloceanspaces.com/uploads/d15f8793-52e9-4c61-8599-db3f9cac1fe5-16.1-background.png",
		"digitalOcean",
		"megustaviajar",
		"uploads/",
		858688,
		"Some title",
		"Some description",
		"Some alt"
	);
	`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestMediaDb_FindByID(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := NewMediaDB(Db)
	media, err := productDb.FindByID("4848560e-b438-4076-8613-4f263d743bf8")
	require.Nil(t, err)
	require.Equal(t, "4848560e-b438-4076-8613-4f263d743bf8", media.GetID())
	require.Equal(t, "0ef22762-de0c-4038-9158-8871fb431aca-16.1-background.png", media.GetName())
	require.Equal(t, "image/png", media.GetContentType())
	require.Equal(t, "https://megustaviajar.sfo2.digitaloceanspaces.com/uploads/d15f8793-52e9-4c61-8599-db3f9cac1fe5-16.1-background.png", media.GetLink())
	require.Equal(t, "digitalOcean", media.GetProvider())
	require.Equal(t, "megustaviajar", media.GetBucketName())
	require.Equal(t, "uploads/", media.GetDirectory())
	require.Equal(t, int64(858688), media.GetSize())
	require.Equal(t, "Some title", media.GetTitle())
	require.Equal(t, "Some description", media.GetDescription())
	require.Equal(t, "Some alt", media.GetAlt())
}

func TestMediaDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	mediaDB := NewMediaDB(Db)

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

	media := entity.MakeMediaEntity(
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

	result, err := mediaDB.Create(media)
	require.Nil(t, err)
	require.Equal(t, result.Title, media.GetTitle())
	require.Equal(t, result.Description, media.GetDescription())
	require.Equal(t, result.Alt, media.GetAlt())
}

func TestMediaDb_FindAll(t *testing.T) {}

func TestMediaDb_Delete(t *testing.T) {}
