package files

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"

	"github.com/h2non/filetype"
)

func Info(r io.ReadSeeker) (*FileInfo, error) {
	fileType, err := filetype.MatchReader(r)
	if err != nil {
		return nil, err
	}

	if fileType == filetype.Unknown {
		return nil, errors.New("unknown file type")
	}

	if err := seekReset(r); err != nil {
		return nil, err
	}

	h := sha256.New()
	fileSize, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	digest := fmt.Sprintf("%x", h.Sum(nil))

	if err := seekReset(r); err != nil {
		return nil, err
	}

	return &FileInfo{
		Extension: fileType.Extension,
		Size:      fileSize,
		Mime: Mime{
			Type:    fileType.MIME.Type,
			Subtype: fileType.MIME.Subtype,
			Value:   fileType.MIME.Value,
		},
		Hash: Hash{
			Algorithm: "sha256",
			Digest:    digest,
		},
	}, nil
}

type Mime struct {
	Type    string
	Subtype string
	Value   string
}

type Hash struct {
	Algorithm string
	Digest    string
}

func (h Hash) String() string {
	return fmt.Sprintf("%s:%s", h.Algorithm, h.Digest)
}

type FileInfo struct {
	Extension string
	Size      int64
	Mime      Mime
	Hash      Hash
}

func seekReset(r io.ReadSeeker) error {
	if pos, err := r.Seek(0, 0); err != nil || pos != 0 {
		return fmt.Errorf("seek reset: %w", err)
	}
	return nil
}
