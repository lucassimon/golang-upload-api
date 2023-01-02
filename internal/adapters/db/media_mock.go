package db

import (
	"github.com/lucassimon/golang-upload-api/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MediaDBRepositoryMock struct {
	mock.Mock
}

func (m *MediaDBRepositoryMock) Create(media *entity.MediaEntity) (*entity.MediaEntity, error) {
	args := m.Called(media)
	// TODO: Corrigir o erro
	return args.Get(0).(*entity.MediaEntity), args.Error(1)
}

func (m *MediaDBRepositoryMock) FindAll(page, limit int, sort string) ([]entity.MediaEntity, error) {
	args := m.Called(page, limit, sort)
	return args.Get(0).([]entity.MediaEntity), args.Error(1)
}

func (m *MediaDBRepositoryMock) FindByID(id string) (*entity.MediaEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.MediaEntity), args.Error(1)
}

func (m *MediaDBRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(1)
}
