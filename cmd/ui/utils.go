package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func decodeJSON(v interface{}, data io.ReadCloser) error {
	if v == nil {
		return fmt.Errorf("decodeJSON::v is nil")
	}

	if data == nil {
		return fmt.Errorf("decodeJSON::body is nil")
	}

	err := json.NewDecoder(data).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
