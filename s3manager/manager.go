package s3manager

import (
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
)

var ErrNotImplemented = errors.New("not implemented")

var _ s3manageriface.DownloaderAPI = (*Manager)(nil)
var _ s3manageriface.UploaderAPI = (*Manager)(nil)
var _ s3manageriface.UploadWithIterator = (*Manager)(nil)
var _ s3manageriface.BatchDelete = (*Manager)(nil)

type Manager struct{}

func (*Manager) Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return 0, ErrNotImplemented
}

func (*Manager) DownloadWithContext(aws.Context, io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return 0, ErrNotImplemented
}

func (*Manager) DownloadWithIterator(aws.Context, s3manager.BatchDownloadIterator, ...func(*s3manager.Downloader)) error {
	return ErrNotImplemented
}

func (*Manager) Upload(*s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return nil, ErrNotImplemented
}

func (*Manager) UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return nil, ErrNotImplemented
}

func (*Manager) UploadWithIterator(aws.Context, s3manager.BatchUploadIterator, ...func(*s3manager.Uploader)) error {
	return ErrNotImplemented
}

func (*Manager) Delete(aws.Context, s3manager.BatchDeleteIterator) error {
	return ErrNotImplemented
}
