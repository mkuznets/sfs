package api

import (
	"context"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/segmentio/ksuid"
	"io"
	"mkuznets.com/go/sps/internal/herror"
	"os"
)

const (
	headerSize int = 512
)

type Uploader interface {
	Upload(ctx context.Context, r io.ReadSeeker) (*UploadInfo, error)
}

type UploadInfo struct {
	Url         string
	Size        int64
	ContentType string
}

type localUploader struct {
}

func NewUploader() Uploader {
	return &localUploader{}
}

func (u *localUploader) Upload(ctx context.Context, r io.ReadSeeker) (*UploadInfo, error) {
	header := make([]byte, headerSize)
	n, err := r.Read(header)
	if err != nil || n != headerSize {
		return nil, herror.Internal("could not read file header").WithError(err)
	}

	fileType, err := filetype.Match(header)
	if err != nil {
		return nil, err
	}
	if fileType == filetype.Unknown || fileType.MIME.Type != "audio" {
		return nil, herror.Validation("unsupported file type")
	}

	tmpFile, err := os.CreateTemp("", fmt.Sprintf("%s_*.%s", ksuid.New().String(), fileType.Extension))
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if pos, err := tmpFile.Seek(0, 0); err != nil || pos != 0 {
		return nil, herror.Internal("could not seek to the beginning of the file").WithError(err)
	}

	fileSize, err := io.Copy(tmpFile, r)
	if _, err := io.Copy(tmpFile, r); err != nil {
		return nil, err
	}

	return &UploadInfo{
		Url:         tmpFile.Name(),
		Size:        fileSize,
		ContentType: fileType.MIME.Value,
	}, nil
}
