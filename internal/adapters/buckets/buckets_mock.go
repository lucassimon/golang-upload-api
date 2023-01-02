package buckets

import (
	"context"
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type ProviderMock struct {
	mock.Mock
}

func (m *ProviderMock) Upload(ctx context.Context, file *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	args := m.Called(ctx, file, uniqueName, extension)

	return args.String(0), args.Error(1)
}
