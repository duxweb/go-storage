package storage

import (
	"context"
	"github.com/duxweb/go-storage/drivers"
	"io"
)

func New(Type string, config map[string]string) FileStorage {
	var driver FileStorage
	switch Type {
	case "local":
		driver = drivers.NewLocalStorage(config["path"], config["domain"])
		break
	case "qiniu":
		driver = drivers.NewQiniuStorage(config["bucket"], config["accessKey"], config["secretKey"], config["domain"])
		break
	case "cos":
		driver = drivers.NewCoStorage(config["secretId"], config["secretKey"], config["region"], config["bucket"], config["domain"])
		break
	case "oss":
		driver = drivers.NewOssStorage(config["accessId"], config["accessSecret"], config["endpoint"], config["bucket"], config["domain"], config["region"])
		break
	case "obs":
		driver = drivers.NewObsStorage(config["ak"], config["sk"], config["endpoint"], config["bucket"], config["domain"])
		break
	}
	return driver
}

type FileStorage interface {
	//Write writes content to a file
	Write(ctx context.Context, path string, contents string, config map[string]any) error
	// WriteStream writes the data stream to a file
	WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]any) error
	// Read a file to a string
	Read(ctx context.Context, path string) (string, error)
	// ReadStream read the file to the stream
	ReadStream(ctx context.Context, path string) (io.Reader, error)
	// Delete delete file
	Delete(ctx context.Context, path string) error
	// PublicUrl gets the file public link
	PublicUrl(ctx context.Context, path string) (string, error)
	//PrivateUrl sets the file private link
	PrivateUrl(ctx context.Context, path string) (string, error)
}
