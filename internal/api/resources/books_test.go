package resources

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

func TestBookList(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/books", nil)
	resp := Record(req, env.BookList)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}

func TestNoIdBookById(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/books", nil)
	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}

/*

// The below test should work, but have not figured out why the book ID
// is not being captured by mux.Vars. TODO.

func TestBookById(t *testing.T) {
	env := NewEnv(nil)

	expected := models.NewBook(
		"1,000 Year Plan",
		"Isaac Asimov",
		"1951",
		"fiction",
		"read",
	)

	err := env.Books.Insert(expected)
	assert.Nil(t, err)

	req := httptest.NewRequest(
		"GET",
		fmt.Sprintf("/books/%s", expected.ID),
		nil,
	)
	assert.Nil(t, err)

	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
*/
