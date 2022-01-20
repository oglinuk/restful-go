package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewBook tests if the value of NewBook is equal to the expected one
func TestNewBook(t *testing.T) {
	expected := &Book{
		"b25228dc32c91679a885204942bec44f",
		"foundations fear",
		"gregory benford",
		"1997",
		"fiction",
		"read",
	}

	actual := NewBook(
		"foundations fear",
		"gregory benford",
		"1997",
		"fiction",
		"read",
	)
	assert.Equal(t, expected, actual)
}
