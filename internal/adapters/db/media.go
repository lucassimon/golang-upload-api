package db

import (
	"database/sql"
	"log"

	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
)

type MediaDBRepositoryInterface interface {
	Create(media *entity.MidiaEntity) error
	FindAll(page, limit int, sort string) ([]entity.MidiaEntity, error)
	FindByID(id string) (*entity.MidiaEntity, error)
	Delete(id string) error
}

type MediaDB struct {
	db *sql.DB
}

func NewMediaDB(db *sql.DB) *MediaDB {
	return &MediaDB{db: db}
}

func (m *MediaDB) Create(media *entity.MidiaEntity) error {
	log.Println("call create from db adapter")

	// stmt, err := m.db.Prepare(`insert into medias(id, name, content_type, link, provider, bucket_name, directory, size) values(?,?,?,?,?,?,?,?)`)
	// if err != nil {
	// 	return err
	// }
	// _, err = stmt.Exec(
	// 	media.GetID(),
	// 	media.GetName(),
	// 	media.GetContentType(),
	// 	media.GetLink(),
	// 	media.GetProvider(),
	// 	media.GetBucketName(),
	// 	media.GetDirectory(),
	// 	media.GetSize(),
	// )
	// if err != nil {
	// 	return err
	// }
	// err = stmt.Close()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (m *MediaDB) FindAll(page, limit int, sort string) ([]entity.MidiaEntity, error) {

	return nil, nil
}

func (m *MediaDB) FindByID(id string) (*entity.MidiaEntity, error) {
	return nil, nil
}

func (m *MediaDB) Delete(id string) error {
	return nil
}
