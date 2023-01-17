package files

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, path string, r io.Reader) (*UploadResult, error)
}

type UploadResult struct {
	Id  string
	Url string
}
