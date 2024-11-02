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
	sign   func(path string) (string, error)
}

func NewLocalStorage(root string, domain string, sign func(path string) (string, error)) *LocalStorage {
	return &LocalStorage{
		root:   root,
		domain: domain,
		sign:   sign,
	}
}

func (s *LocalStorage) Write(ctx context.Context, path string, contents string) error {
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

func (s *LocalStorage) WriteStream(ctx context.Context, path string, stream io.Reader) error {
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

func (s *LocalStorage) Size(ctx context.Context, path string) (int64, error) {
	fullPath := s.root + "/" + path
	stat, err := os.Stat(fullPath)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (s *LocalStorage) Exists(ctx context.Context, path string) (bool, error) {
	fullPath := s.root + "/" + path
	_, err := os.Stat(fullPath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *LocalStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	srcUrl := fmt.Sprintf("%s/%s", strings.TrimRight(s.domain, "/"), path)
	srcUri, _ := url.Parse(srcUrl)
	return srcUri.String(), nil
}

func (s *LocalStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	return s.domain + "/" + path, nil
}

func (s *LocalStorage) SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error) {
	sign, err := s.getSign(path)
	if err != nil {
		return url, nil, err
	}
	return url, map[string]string{
		"sign": sign,
	}, nil
}

func (s *LocalStorage) SignPutUrl(ctx context.Context, path string) (url string, err error) {
	sign, err := s.getSign(path)
	if err != nil {
		return url, err
	}
	if strings.Contains(url, "?") {
		url = url + "&sign=" + sign
	} else {
		url = url + "?sign=" + sign
	}
	return url, nil
}

func (s *LocalStorage) getSign(path string) (string, error) {
	if s.sign != nil {
		sign, err := s.sign(path)
		if err != nil {
			return "", err
		}
		return sign, nil
	}
	return "", nil
}
