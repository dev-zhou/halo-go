package _decode

import (
	"encoding/json"
	"io"
)

func DecodeJSON[R comparable](content io.Reader, result R) error {
	decoder := json.NewDecoder(content)
	if err := decoder.Decode(&result); err != nil {
		return err
	}
	return nil
}
