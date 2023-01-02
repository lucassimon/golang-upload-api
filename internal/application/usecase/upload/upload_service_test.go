package upload

// func TestUploadService_Save(t *testing.T) {
// 	name := "Foo"
// 	contentType := "image/png"
// 	link := "http://some-url.com"
// 	provider := "amazon"
// 	bucket_name := "Foo"
// 	directory := "/uploads/some-file.png"
// 	var size int64 = 8045687521
// 	title := "Some Title Image"
// 	description := "SOme description"
// 	alt := "Some alternative"

// 	media := entity.MakeMediaEntity(
// 		name,
// 		contentType,
// 		link,
// 		provider,
// 		bucket_name,
// 		directory,
// 		size,
// 		title,
// 		description,
// 		alt,
// 	)

// 	ctx := context.Background()

// 	providerMock := &buckets.ProviderMock{}
// 	bucket := &buckets.BucketFactory{
// 		Name:       "test",
// 		BucketName: "some-bucket-name/",
// 		Path:       buckets.UPLOADS_PREFIX,
// 		Provider:   providerMock,
// 	}

// 	repositoryMock := &db.MediaDBRepositoryMock{}
// 	repositoryMock.On("Create", media).Return(nil)

// 	service := NewUploadService(bucket, repositoryMock)
// 	err := service.Save(ctx, media)

// 	assert.Nil(t, err)

// 	repositoryMock.AssertExpectations(t)
// }
