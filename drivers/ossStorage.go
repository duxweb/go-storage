package drivers

import (
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"io"
	"net/url"
	"strings"
)

type OssStorage struct {
	Client     *oss.Client
	Domain     string
	BucketName string
}

func NewOssStorage(accessKeyId, accessKeySecret, endpoint, bucketName, domain, region string) *OssStorage {

	provider := credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret)
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithEndpoint(endpoint).WithRegion(region)
	client := oss.NewClient(cfg)

	return &OssStorage{
		Client:     client,
		Domain:     domain,
		BucketName: bucketName,
	}
}

func (ofs *OssStorage) Write(ctx context.Context, path string, contents string, config map[string]interface{}) error {
	return ofs.WriteStream(ctx, path, strings.NewReader(contents), config)
}

func (ofs *OssStorage) WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]interface{}) error {
	u := ofs.Client.NewUploader()
	_, err := u.UploadFrom(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(ofs.BucketName),
		Key:    oss.Ptr(path),
	}, stream)
	if err != nil {
		return err
	}
	return nil
}

func (ofs *OssStorage) Read(ctx context.Context, path string) (string, error) {
	stream, err := ofs.ReadStream(ctx, path)
	if err != nil {
		return "", err
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, stream)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (ofs *OssStorage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	res, err := ofs.Client.GetObject(ctx, &oss.GetObjectRequest{
		Bucket: oss.Ptr(ofs.BucketName),
		Key:    oss.Ptr(path),
	})
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (ofs *OssStorage) Delete(ctx context.Context, path string) error {
	_, err := ofs.Client.DeleteObject(ctx, &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(ofs.BucketName),
		Key:    oss.Ptr(path),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ofs *OssStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	srcUrl := fmt.Sprintf("%s/%s", strings.TrimRight(ofs.Domain, "/"), path)
	srcUri, _ := url.Parse(srcUrl)
	finalUrl := srcUri.String()
	return finalUrl, nil
}

func (ofs *OssStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	res, err := ofs.Client.Presign(ctx, &oss.GetObjectRequest{
		Bucket: oss.Ptr(ofs.BucketName),
		Key:    oss.Ptr(path),
	},
		oss.PresignExpires(3600))
	if err != nil {
		return "", err
	}
	return res.URL, nil
}
