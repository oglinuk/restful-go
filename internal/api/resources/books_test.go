package resources

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
	"github.com/oglinuk/restful-go/internal/pkg/config"
	"github.com/oglinuk/restful-go/internal/pkg/models"
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

func TestGetBookById(t *testing.T) {
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

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", expected.ID)
	req = req.WithContext(context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
	))

	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}

func TestDeleteBookById(t *testing.T) {
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
		"DELETE",
		fmt.Sprintf("/books/%s", expected.ID),
		nil,
	)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", expected.ID)
	req = req.WithContext(context.WithValue(
			req.Context(),
			chi.RouteCtxKey,
			rctx,
	))

	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
