package db

import (
	"database/sql"
	"log"
	"strings"

	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type MediaDBRepositoryInterface interface {
	Create(media *entity.MediaEntity) (*entity.MediaEntity, error)
	FindAll(page, limit int, sort string) ([]entity.MediaEntity, error)
	FindByID(id string) (*entity.MediaEntity, error)
	Delete(id string) error
}

type MediaDB struct {
	db *sql.DB
}

func NewMediaDB(db *sql.DB) *MediaDB {
	return &MediaDB{db: db}
}

func (m *MediaDB) Create(media *entity.MediaEntity) (*entity.MediaEntity, error) {
	log.Println("call create from db adapter")

	stmt, err := m.db.Prepare(`INSERT INTO medias(id, name, content_type, link, provider, bucket_name, directory, size, title, description, alt) VALUES(?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		media.GetID(),
		media.GetName(),
		media.GetContentType(),
		media.GetLink(),
		media.GetProvider(),
		media.GetBucketName(),
		media.GetDirectory(),
		media.GetSize(),
		media.GetTitle(),
		media.GetDescription(),
		media.GetAlt(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (m *MediaDB) FindAll(page, limit int, sort string) ([]entity.MediaEntity, error) {
	var medias []entity.MediaEntity
	var err error

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	sort_lower := strings.ToLower(sort)

	if sort_lower != "" && sort_lower != "asc" && sort_lower != "desc" {
		sort_lower = "asc"
	}

	rows, err := m.db.Query("SELECT * FROM medias ORDER BY id ? LIMIT ? OFFSET ?", sort_lower, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var media entity.MediaEntity
		err = rows.Scan(&media.Id, &media.Name, &media.ContentType, &media.Link, &media.Provider, &media.BucketName, &media.Directory, &media.Size, &media.Title, &media.Description, &media.Alt)
		if err != nil {
			return nil, err
		}
		medias = append(medias, media)
	}
	return medias, err
}

func (m *MediaDB) FindByID(id string) (*entity.MediaEntity, error) {
	var media entity.MediaEntity
	stmt, err := m.db.Prepare("SELECT * from medias where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&media.Id, &media.Name, &media.ContentType, &media.Link, &media.Provider, &media.BucketName, &media.Directory, &media.Size, &media.Title, &media.Description, &media.Alt)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (m *MediaDB) Delete(id string) error {
	stmt, err := m.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
