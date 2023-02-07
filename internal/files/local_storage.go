package files

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

type localStorage struct {
	basePath  string
	urlPrefix string
}

func NewLocalStorage(basePath string, urlPrefix string) Storage {
	return &localStorage{
		basePath:  basePath,
		urlPrefix: urlPrefix,
	}
}

func (l *localStorage) ensurePath(path string) (string, error) {
	absPath, err := filepath.Abs(filepath.Join(l.basePath, path))
	if err != nil {
		return "", fmt.Errorf("get absolute path for %s: %w", absPath, err)
	}

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return "", fmt.Errorf("create target directory: %w", err)
	}

	return absPath, nil
}

func (l *localStorage) Upload(_ context.Context, path string, r io.Reader) (*UploadResult, error) {
	absPath, err := l.ensurePath(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(absPath)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer func() {
		_ = f.Sync()
		_ = f.Close()
	}()

	if _, err := io.Copy(f, r); err != nil {
		return nil, fmt.Errorf("copy file: %w", err)
	}

	u, err := url.Parse(l.urlPrefix)
	u = u.JoinPath(path)

	return &UploadResult{
		Id:  absPath,
		Url: u.String(),
	}, nil
}
