package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGet tests if the value from Get is the same as the defaultCfg
func TestGet(t *testing.T) {
	expected := defaultCfg
	actual := Get()
	assert.Equal(t, actual, expected)

	t.Cleanup(func() {
		os.Remove(actual.Name)
	})
}
