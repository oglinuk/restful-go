package api

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

// TestRun calls Init in a goroutine, waits 42 milliseconds, then makes
// a request to each resource endpoint.
func TestRun(t *testing.T) {
	go Run()
	time.Sleep(time.Millisecond*42)

	cfg := config.Get()
	addr := fmt.Sprintf("http://%s:%d", cfg.Server.Host, cfg.Server.Port)


	//TODO: Replace below with loop over cfg.Server.Endpoints
	_, err := http.Get(addr)
	assert.Nil(t, err)

	_, err = http.Get(fmt.Sprintf("%s/books", addr))
	assert.Nil(t, err)

	t.Cleanup(func() {
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
