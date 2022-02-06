package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNilInputsDecodeJSON ensures that nil inputs are handled
func TestNilInputsDecodeJSON(t *testing.T) {
	emptyResp := &http.Response{}
	expectedNilInterfaceErr := fmt.Errorf("decodeJSON::v is nil")

	actualNilInterfaceErr := decodeJSON(nil, emptyResp.Body)
	assert.Equal(t, expectedNilInterfaceErr, actualNilInterfaceErr)

	emptyInterface := map[string]string{}
	expectedNilBodyErr := fmt.Errorf("decodeJSON::body is nil")

	actualNilBodyErr := decodeJSON(emptyInterface, nil)
	assert.Equal(t, expectedNilBodyErr, actualNilBodyErr)
}

// TestEOFDecodeJSON ensures that EOF errors are handled
func TestEOFDecodeJSON(t *testing.T) {
	emptyInterface := map[string]string{}
	eofReq := &http.Response{
		Body: io.NopCloser(bytes.NewBufferString(`"{'test': 'eof}`)),
	}

	actualEOFBodyErr := decodeJSON(emptyInterface, eofReq.Body)
	assert.Equal(t, io.ErrUnexpectedEOF, actualEOFBodyErr)
}

// TestDecodeJSON ensures valid JSON is handled
func TestDecodeJSON(t *testing.T) {
	expected := map[string]string{"hello": "world!"}

	actual := make(map[string]string)

	validJSON := bytes.NewBufferString(`{"hello": "world!"}`)
	err := decodeJSON(&actual, io.NopCloser(validJSON))
	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}
