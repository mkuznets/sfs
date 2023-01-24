package yrender

import (
	"encoding/json"
	"io"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
)

func DecodeJson[T any](r io.Reader) (T, error) {
	var v T
	defer func(src io.Reader) {
		_, _ = io.Copy(io.Discard, src)
	}(r)
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, yerr.Invalid("invalid JSON request").WithCause(err)
	}
	return v, nil
}
