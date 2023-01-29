package yjson

import (
	"encoding/json"
	"io"
	"mkuznets.com/go/sfs/ytils/yerr"
)

func Decode[T any](r io.Reader) (T, error) {
	var v T
	defer func(src io.Reader) {
		_, _ = io.Copy(io.Discard, src)
	}(r)
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, yerr.Invalid("invalid JSON").Err(err)
	}
	return v, nil
}

func Unmarshall[T any](data []byte) (*T, error) {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return &v, yerr.Invalid("invalid JSON").Err(err)
	}
	return &v, nil
}
