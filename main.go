package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"io"
	"net/url"
	"strings"
	"time"
)

func New(configMap map[string]string) (*Store, error) {

	cred := credentials.NewStaticCredentialsProvider(configMap["accessKey"], configMap["secretKey"], "")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(cred),
		config.WithRegion(configMap["region"]),
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

	return &Store{
		s3:     client,
		config: configMap,
	}, nil
}

func getUrl(config map[string]string) string {
	prefix := "https://"
	if config["ssl"] != "" {
		prefix = "http://"
	}
	return prefix + config["endpoint"]
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

type Store struct {
	s3     *s3.Client
	config map[string]string
}

func (s *Store) Write(ctx context.Context, path string, contents string) error {
	render := strings.NewReader(contents)
	return s.WriteStream(ctx, path, render)
}

func (s *Store) WriteStream(ctx context.Context, path string, stream io.Reader) error {
	input := &s3.PutObjectInput{
		Body:   stream,
		Bucket: aws.String(s.config["bucket"]),
		Key:    aws.String(path),
	}
	_, err := s.s3.PutObject(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Read(ctx context.Context, path string) (string, error) {
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

func (s *Store) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config["bucket"]),
		Key:    aws.String(path),
	}

	out, err := s.s3.GetObject(ctx, input)
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func (s *Store) Delete(ctx context.Context, path string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.config["bucket"]),
		Key:    aws.String(path),
	}

	_, err := s.s3.DeleteObject(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) PublicUrl(ctx context.Context, path string) (string, error) {
	srcUrl := fmt.Sprintf("%s/%s", strings.TrimRight(s.config["domain"], "/"), path)
	srcUri, _ := url.Parse(srcUrl)
	finalUrl := srcUri.String()
	return finalUrl, nil
}

func (s *Store) PrivateUrl(ctx context.Context, path string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config["bucket"]),
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

func (s *Store) SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.config["bucket"]),
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

func (s *Store) SignPutUrl(ctx context.Context, path string) (url string, err error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.config["bucket"]),
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
