package s3manager

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"storj.io/s3link/s3link"
)

var _ s3manageriface.DownloaderAPI = (*Manager)(nil)
var _ s3manageriface.UploaderAPI = (*Manager)(nil)
var _ s3manageriface.UploadWithIterator = (*Manager)(nil)
var _ s3manageriface.BatchDelete = (*Manager)(nil)

type Manager struct{}

func (_ *Manager) Download(io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return 0, s3link.NotImplemented
}

func (_ *Manager) DownloadWithContext(aws.Context, io.WriterAt, *s3.GetObjectInput, ...func(*s3manager.Downloader)) (int64, error) {
	return 0, s3link.NotImplemented
}

func (_ *Manager) DownloadWithIterator(aws.Context, s3manager.BatchDownloadIterator, ...func(*s3manager.Downloader)) error {
	return s3link.NotImplemented
}

func (_ *Manager) Upload(*s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return nil, s3link.NotImplemented
}

func (_ *Manager) UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return nil, s3link.NotImplemented
}

func (_ *Manager) UploadWithIterator(aws.Context, s3manager.BatchUploadIterator, ...func(*s3manager.Uploader)) error {
	return s3link.NotImplemented
}

func (_ *Manager) Delete(aws.Context, s3manager.BatchDeleteIterator) error {
	return s3link.NotImplemented
}
