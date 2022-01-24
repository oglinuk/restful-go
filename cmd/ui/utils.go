package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func decodeJSON(v interface{}, body io.ReadCloser) error {
	if v == nil {
		return fmt.Errorf("decodeJSON::v is nil")
	}

	if body == nil {
		return fmt.Errorf("decodeJSON::body is nil")
	}
	defer body.Close()

	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
