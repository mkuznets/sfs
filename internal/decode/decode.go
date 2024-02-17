package decode

import (
	"encoding/json"
	"fmt"
	"io"
)

func JSON[T any](r io.Reader) (*T, error) {
	var v T
	defer io.Copy(io.Discard, r)

	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return nil, fmt.Errorf("decode JSON: %w", err)
	}
	return &v, nil
}
