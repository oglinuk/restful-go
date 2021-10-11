package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewBook tests if the value of NewBook is equal to the expected one
func TestNewBook(t *testing.T) {
	expected := &Book{
		"75f10dcf2b3bb41eff76fa50e664d132",
		"foundations fear",
		"gregory benford",
		"946684800",
	}
	actual := NewBook("foundations fear", "gregory benford", "946684800")
	assert.Equal(t, expected, actual)
}
