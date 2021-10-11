package resources

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeartbeat(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/", nil)
	resp := Record(req, env.Heartbeat)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
