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
	Root   string
	Domain string
	Path   string
	Sign   func(path string) (string, error)
}

func NewLocalStorage(configMap map[string]string, signs ...func(path string) (string, error)) *LocalStorage {

	store := &LocalStorage{
		Root:   configMap["root"],
		Domain: configMap["domain"],
		Path:   configMap["path"],
	}
	if len(signs) > 0 {
		store.Sign = signs[0]
	}
	return store
}

func (s *LocalStorage) Write(ctx context.Context, path string, contents string, metadata ...map[string]string) error {
	fullPath := s.Root + "/" + s.getUploadPath(path)
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

func (s *LocalStorage) WriteStream(ctx context.Context, path string, stream io.Reader, metadata ...map[string]string) error {
	fullPath := s.Root + "/" + s.getUploadPath(path)
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
	fullPath := s.Root + "/" + s.getUploadPath(path)
	contents, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func (s *LocalStorage) ReadStream(ctx context.Context, path string) (io.Reader, error) {
	fullPath := s.Root + "/" + s.getUploadPath(path)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *LocalStorage) Delete(ctx context.Context, path string) error {
	fullPath := s.Root + "/" + s.getUploadPath(path)
	return os.Remove(fullPath)
}

func (s *LocalStorage) Size(ctx context.Context, path string) (int64, error) {
	fullPath := s.Root + "/" + s.getUploadPath(path)
	stat, err := os.Stat(fullPath)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (s *LocalStorage) Exists(ctx context.Context, path string) (bool, error) {
	fullPath := s.Root + "/" + s.getUploadPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *LocalStorage) PublicUrl(ctx context.Context, path string) (string, error) {
	domain := strings.TrimRight(s.Domain, "/")
	if s.Path != "" {
		domain = fmt.Sprintf("%s/%s", s.Domain, s.Path)
	}
	srcUrl := fmt.Sprintf("%s/%s", domain, path)
	srcUri, _ := url.Parse(srcUrl)
	return srcUri.String(), nil
}

func (s *LocalStorage) PrivateUrl(ctx context.Context, path string) (string, error) {
	return s.PublicUrl(ctx, path)
}

func (s *LocalStorage) SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error) {
	url = s.getUploadPath(path)

	sign, err := s.getSign(url)
	if err != nil {
		return url, nil, err
	}

	return url, map[string]string{
		"sign": sign,
		"key":  path,
	}, nil
}

func (s *LocalStorage) SignPutUrl(ctx context.Context, path string) (string, error) {
	u := s.getUploadPath(path)

	sign, err := s.getSign(u)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("sign", sign)
	params.Add("key", path)

	q := params.Encode()

	if strings.Contains(u, "?") {
		u = u + "&sign=" + q
	} else {
		u = u + "?sign=" + q
	}
	return u, nil
}

func (s *LocalStorage) getUploadPath(path string) string {
	if s.Path != "" {
		path = fmt.Sprintf("%s/%s", s.Path, path)
	}
	return path
}

func (s *LocalStorage) getSign(path string) (string, error) {
	if s.Sign != nil {
		sign, err := s.Sign(path)
		if err != nil {
			return "", err
		}
		return sign, nil
	}
	return "", nil
}

func (s *LocalStorage) Local() bool {
	return true
}
