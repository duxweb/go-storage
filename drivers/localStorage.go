package drivers

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type LocalStorage struct {
	root   string
	domain string
}

func NewLocalStorage(root string, domain string) *LocalStorage {
	return &LocalStorage{root, domain}
}

func (s *LocalStorage) Write(ctx context.Context, path string, contents string, config map[string]interface{}) error {
	fullPath := s.root + "/" + path
	paths, _ := filepath.Split(fullPath)
	err := os.MkdirAll(paths, 0777)
	if err != nil {
		return err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocalStorage) WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]interface{}) error {
	fullPath := s.root + "/" + path
	paths, _ := filepath.Split(fullPath)
	err := os.MkdirAll(paths, 0777)
	if err != nil {
		return err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, stream)
	if err != nil {
		return err
	}
	return nil
}

func (s *LocalStorage) Read(ctx context.Context, path string) (string, error) {
	fullPath := s.root + "/" + path
	contents, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func (s *LocalStorage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	fullPath := s.root + "/" + path
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *LocalStorage) Delete(ctx context.Context, path string) error {
	fullPath := s.root + "/" + path
	return os.Remove(fullPath)
}

func (s *LocalStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	srcUrl := fmt.Sprintf("%s/%s", strings.TrimRight(s.domain, "/"), path)
	srcUri, _ := url.Parse(srcUrl)
	return srcUri.String(), nil
}

func (s *LocalStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	return s.domain + "/" + path, nil
}
