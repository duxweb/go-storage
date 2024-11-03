package drivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"io"
	"net/url"
	"strings"
	"time"
)

type S3Storage struct {
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
	Domain    string
	Bucket    string
	SSL       bool
	Immutable bool
	s3        *s3.Client
}

func NewS3Storage(configMap map[string]string) (*S3Storage, error) {

	store := S3Storage{
		Endpoint:  configMap["endpoint"],
		Region:    configMap["region"],
		AccessKey: configMap["accessKey"],
		SecretKey: configMap["secretKey"],
		Bucket:    configMap["bucket"],
		Domain:    configMap["domain"],
		SSL:       configMap["ssl"] == "true",
		Immutable: configMap["immutable"] == "true",
	}

	cred := credentials.NewStaticCredentialsProvider(store.AccessKey, store.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(cred),
		config.WithRegion(store.Region),
	)
	if err != nil {
		return nil, err
	}

	var client *s3.Client

	if configMap["immutable"] == "true" {
		client = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.EndpointResolverV2 = &staticResolver{
				url: getUrl(configMap),
			}
		})
	} else {
		client = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(getUrl(configMap))
		})
	}

	store.s3 = client

	return &store, nil
}

func (s *S3Storage) Write(ctx context.Context, path string, contents string) error {
	render := strings.NewReader(contents)
	return s.WriteStream(ctx, path, render)
}

func (s *S3Storage) WriteStream(ctx context.Context, path string, stream io.Reader) error {
	input := &s3.PutObjectInput{
		Body:   stream,
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}
	_, err := s.s3.PutObject(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *S3Storage) Read(ctx context.Context, path string) (string, error) {
	stream, err := s.ReadStream(ctx, path)
	if err != nil {
		return "", err
	}
	str, err := io.ReadAll(stream)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func (s *S3Storage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}

	out, err := s.s3.GetObject(ctx, input)
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func (s *S3Storage) Delete(ctx context.Context, path string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}

	_, err := s.s3.DeleteObject(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *S3Storage) head(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	}
	return s.s3.HeadObject(ctx, input)
}

func (s *S3Storage) Size(ctx context.Context, path string) (int64, error) {
	output, err := s.head(ctx, path)
	if err != nil {
		return 0, err
	}
	return *output.ContentLength, nil
}

func (s *S3Storage) Exists(ctx context.Context, path string) (bool, error) {
	_, err := s.head(ctx, path)
	if err != nil {
		var e *smithy.OperationError
		if errors.As(err, &e) {
			if strings.Contains(e.Err.Error(), "404") {
				return false, nil
			}
		}
		return false, err
	}
	return true, err
}

func (s *S3Storage) PublicUrl(ctx context.Context, path string) (string, error) {
	srcUrl := fmt.Sprintf("%s/%s", strings.TrimRight(s.Domain, "/"), path)
	srcUri, _ := url.Parse(srcUrl)
	finalUrl := srcUri.String()
	return finalUrl, nil
}

func (s *S3Storage) PrivateUrl(ctx context.Context, path string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}
	svc := s3.NewPresignClient(s.s3)
	req, err := svc.PresignGetObject(ctx, input, func(o *s3.PresignOptions) {
		o.Expires = 20 * time.Minute
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *S3Storage) SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}
	svc := s3.NewPresignClient(s.s3)
	req, err := svc.PresignPostObject(ctx, input, func(o *s3.PresignPostOptions) {
		o.Expires = 20 * time.Minute
	})
	if err != nil {
		return "", nil, err
	}
	return req.URL, req.Values, nil
}

func (s *S3Storage) SignPutUrl(ctx context.Context, path string) (url string, err error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}
	svc := s3.NewPresignClient(s.s3)
	req, err := svc.PresignPutObject(ctx, input, func(o *s3.PresignOptions) {
		o.Expires = 20 * time.Minute
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *S3Storage) Local() bool {
	return false
}

type staticResolver struct {
	url string
}

func (t *staticResolver) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {

	u, err := url.Parse(t.url)
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{
		URI: *u,
	}, nil
}

func getUrl(config map[string]string) string {
	prefix := "https://"
	if config["ssl"] != "" {
		prefix = "http://"
	}
	return prefix + config["endpoint"]
}
