package router

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New(nil))

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
