package drivers

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
	"strings"
)

type ObsStorage struct {
	client     *obs.ObsClient
	bucketName string
	publicUrl  string
}

func NewObsStorage(ak string, sk string, endpoint string, bucketName string, publicUrl string) *ObsStorage {
	client, err := obs.New(ak, sk, endpoint)
	if err != nil {
		return nil
	}
	return &ObsStorage{
		client:     client,
		bucketName: bucketName,
		publicUrl:  publicUrl,
	}
}

func (h *ObsStorage) Write(ctx context.Context, path string, contents string, config map[string]interface{}) error {
	return h.WriteStream(ctx, path, strings.NewReader(contents), config)
}

func (h *ObsStorage) WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]interface{}) error {
	input := &obs.PutObjectInput{}
	input.Bucket = h.bucketName
	input.Key = path
	input.Body = stream
	_, err := h.client.PutObject(input)
	return err
}

func (h *ObsStorage) Read(ctx context.Context, path string) (string, error) {
	stream, err := h.ReadStream(ctx, path)
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

func (h *ObsStorage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = h.bucketName
	input.Key = path
	output, err := h.client.GetObject(input)
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func (h *ObsStorage) Delete(ctx context.Context, path string) error {
	input := &obs.DeleteObjectInput{}
	input.Bucket = h.bucketName
	input.Key = path
	_, err := h.client.DeleteObject(input)
	return err
}

func (h *ObsStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	return h.publicUrl + "/" + path, nil
}

func (h *ObsStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	input := &obs.CreateSignedUrlInput{
		Method:  obs.HttpMethodGet,
		Bucket:  h.bucketName,
		Key:     path,
		Expires: 3600,
	}
	getObjectOutput, err := h.client.CreateSignedUrl(input)
	if err != nil {
		return "", err
	}
	return getObjectOutput.SignedUrl, nil
}
