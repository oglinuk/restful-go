package resources

import (
	"os"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

func TestHeartbeat(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/", nil)
	resp := Record(req, env.Heartbeat)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
