package files

import (
	"context"
	"io"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
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
		return "", yerr.New("failed to get absolute path for %s", absPath).Err(err)
	}

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return "", yerr.New("failed to create target directory").Err(err)
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
		return nil, yerr.New("failed to create file").Err(err)
	}
	defer func() {
		_ = f.Sync()
		_ = f.Close()
	}()

	if _, err := io.Copy(f, r); err != nil {
		return nil, yerr.New("failed to copy file").Err(err)
	}

	u, err := url.Parse(l.urlPrefix)
	u = u.JoinPath(path)

	return &UploadResult{
		Id:  absPath,
		Url: u.String(),
	}, nil
}
