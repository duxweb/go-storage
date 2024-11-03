package storage

import (
	"context"
	"github.com/duxweb/go-storage/v2/drivers"
	"io"
)

func New(Type string, config map[string]string, signs ...func(path string) (string, error)) (FileStorage, error) {
	var driver FileStorage
	var err error
	switch Type {
	case "local":
		driver = drivers.NewLocalStorage(config, signs...)
		break
	case "s3":
		driver, err = drivers.NewS3Storage(config)
		if err != nil {
			return nil, err
		}
		break
	}
	return driver, nil
}

type FileStorage interface {
	//Write writes content to a file
	Write(ctx context.Context, path string, contents string) error
	// WriteStream writes the data stream to a file
	WriteStream(ctx context.Context, path string, stream io.Reader) error
	// Read a file to a string
	Read(ctx context.Context, path string) (string, error)
	// ReadStream read the file to the stream
	ReadStream(ctx context.Context, path string) (io.Reader, error)
	// Delete delete file
	Delete(ctx context.Context, path string) error
	// Size get file size
	Size(ctx context.Context, path string) (int64, error)
	// Exists checks if file exists
	Exists(ctx context.Context, path string) (bool, error)
	// PublicUrl gets the file public link
	PublicUrl(ctx context.Context, path string) (string, error)
	//PrivateUrl sets the file private link
	PrivateUrl(ctx context.Context, path string) (string, error)
	// SignPostUrl Signature POST Upload Url
	SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error)
	// SignPutUrl Signature PUT Upload Url
	SignPutUrl(ctx context.Context, path string) (url string, err error)
	// Local drive or not
	Local() bool
}
