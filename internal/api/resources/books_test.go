package resources

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookList(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/books", nil)
	resp := Record(req, env.BookList)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
