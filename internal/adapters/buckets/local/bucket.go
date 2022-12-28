package local

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

// Progress is used to track the progress of a file upload.
// It implements the io.Writer interface so it can be passed
// to an io.TeeReader()
type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

type ProviderLocal struct {
}

func NewLocalBucket(ctx context.Context) *ProviderLocal {
	return &ProviderLocal{}
}

func (b *ProviderLocal) Upload(ctx context.Context, fh *multipart.FileHeader, uniqueName string, extension string) (string, error) {
	log.Printf("uploading the file: %s in bucket Local", uniqueName)
	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		return "", err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("./media/uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	directory := fmt.Sprintf("uploads/%s", uniqueName)

	f, err := os.Create(fmt.Sprintf("./media/%s", directory))
	if err != nil {
		return "", err
	}
	defer f.Close()

	pr := &Progress{
		TotalSize: fh.Size,
	}

	_, err = io.Copy(f, io.TeeReader(file, pr))
	if err != nil {
		return "", err
	}

	return directory, nil
}
