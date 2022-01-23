package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
	"github.com/oglinuk/restful-go/internal/pkg/models"
)

func TestCreateBook(t *testing.T) {
	env := NewEnv(nil)

	expected := models.NewBook(
		"Chaos: Making a New Science",
		"James Gleick",
		"1988",
		"non-fiction",
		"reading",
	)

	data, err := json.Marshal(expected)
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(data))
	resp := Record(req, env.CreateBook)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBookList(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/books", nil)
	resp := Record(req, env.BookList)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestNoIdBookById(t *testing.T) {
	env := NewEnv(nil)
	req := httptest.NewRequest("GET", "/books", nil)
	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
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

	req = ChiURLParams(map[string]string{"id": expected.ID}, req)
	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateBookById(t *testing.T) {
	env := NewEnv(nil)

	book := models.NewBook(
		"The Time Machine",
		"H. G. Wells",
		"1979",
		"fiction",
		"reading",
	)

	err := env.Books.Insert(book)
	assert.Nil(t, err)

	updated := models.NewBook(
		"The Time Machine",
		"H. G. Wells",
		"1979",
		"fiction",
		"read",
	)

	id, err := env.Books.Update(book.ID, updated)
	assert.Nil(t, err)
	assert.Equal(t, updated.ID, id)
}

func TestDeleteBookById(t *testing.T) {
	env := NewEnv(nil)

	expected := models.NewBook(
		"The End of Eternity",
		"Isaac Asimov",
		"1955",
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

	req = ChiURLParams(map[string]string{"id": expected.ID}, req)
	resp := Record(req, env.BookById)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
