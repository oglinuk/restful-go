package router

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

func TestNewRouter(t *testing.T) {
	assert.NotNil(t, NewRouter(nil))
}

func TestRouterRoutes(t *testing.T) {
	r := NewRouter(nil)
	assert.NotNil(t, r)

	expectedPatterns := []string{"/", "/books/*"}

	for i, route := range r.Routes() {
		assert.Equal(t, expectedPatterns[i], route.Pattern)
	}

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
