package filestore

import (
	"context"
	"io"
)

type FileStore interface {
	Upload(ctx context.Context, path string, body io.Reader) (*UploadResult, error)
}

type UploadResult struct {
	ID  string
	URL string
}
