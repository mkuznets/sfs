package yrender

import (
	"encoding/json"
	"io"
)

func DecodeJson(r io.Reader, v interface{}) error {
	defer func(src io.Reader) {
		_, _ = io.Copy(io.Discard, src)
	}(r)
	return json.NewDecoder(r).Decode(v)
}
