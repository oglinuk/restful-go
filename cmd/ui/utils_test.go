package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNilInputsDecodeJSON ensures that decodeJSON handles nil inputs
func TestNilInputsDecodeJSON(t *testing.T) {
	badReq := &http.Response{}
	expectedNilInterfaceErr := fmt.Errorf("decodeJSON::v is nil")

	actualNilInterfaceErr := decodeJSON(nil, badReq.Body)
	assert.NotNil(t, actualNilInterfaceErr)
	assert.Equal(t, expectedNilInterfaceErr, actualNilInterfaceErr)

	badInterface := map[string]string{}
	expectedNilBodyErr := fmt.Errorf("decodeJSON::body is nil")

	actualNilBodyErr := decodeJSON(badInterface, nil)
	assert.NotNil(t, actualNilBodyErr)
	assert.Equal(t, expectedNilBodyErr, actualNilBodyErr)

	badReq.Body = io.NopCloser(bytes.NewBufferString(`"{'test': 'bad}`))

	actualNilBodyDecoderErr := decodeJSON(badInterface, badReq.Body)
	assert.NotNil(t, actualNilBodyDecoderErr)
	assert.Equal(t, io.ErrUnexpectedEOF, actualNilBodyDecoderErr)
}
