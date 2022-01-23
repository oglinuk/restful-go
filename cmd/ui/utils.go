package main

import (
	"encoding/json"
	"io"
)

func decodeJSON(v interface{}, body io.ReadCloser) error {
	defer body.Close()

	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
