package drivers

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"strings"
)

type QiniuFileStorage struct {
	Bucket string
	Domain string
	Mac    *qbox.Mac
}

func NewQiniuStorage(Bucket, AccessKey, SecretKey, Domain string) *QiniuFileStorage {
	return &QiniuFileStorage{
		Mac:    qbox.NewMac(AccessKey, SecretKey),
		Domain: Domain,
		Bucket: Bucket,
	}
}

func (qfs *QiniuFileStorage) Write(ctx context.Context, path string, contents string, config map[string]any) error {
	return qfs.WriteStream(ctx, path, strings.NewReader(contents), config)
}

func (qfs *QiniuFileStorage) WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]any) error {
	putPolicy := storage.PutPolicy{
		Scope: qfs.Bucket + ":" + path,
	}
	upToken := putPolicy.UploadToken(qfs.Mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	params := map[string]string{}
	for key, value := range config {
		params[key] = value.(string)
	}
	putExtra := storage.PutExtra{
		Params: params,
	}
	err := formUploader.Put(ctx, &ret, upToken, path, stream, -1, &putExtra)
	if err != nil {
		return err
	}
	return nil
}

func (qfs *QiniuFileStorage) Read(ctx context.Context, path string) (string, error) {
	stream, err := qfs.ReadStream(ctx, path)
	if err != nil {
		return "", err
	}
	str, err := io.ReadAll(stream)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func (qfs *QiniuFileStorage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	url, err := qfs.PrivateUrl(ctx, path)
	if err != nil {
		return nil, err
	}
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to read file " + path)
	}
	return strings.NewReader(resp.String()), nil
}

func (qfs *QiniuFileStorage) Delete(ctx context.Context, path string) error {
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(qfs.Mac, &cfg)
	return bucketManager.Delete(qfs.Bucket, path)
}

func (qfs *QiniuFileStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	return storage.MakePublicURL(qfs.Domain, path), nil
}

func (qfs *QiniuFileStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	return storage.MakePrivateURL(qfs.Mac, qfs.Domain, path, 3600), nil
}
