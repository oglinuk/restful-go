package repositories

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
	"github.com/oglinuk/restful-go/internal/pkg/database"
	"github.com/oglinuk/restful-go/internal/pkg/models"
)

func TestNewBooksRepo(t *testing.T) {
	br := NewBooksRepo(nil)
	assert.NotNil(t, br)
}

func TestInsertBook(t *testing.T) {
	br := NewBooksRepo(database.Open(bookSchema))

	err := br.Insert(models.NewBook(
		"1,000 Year Plan",
		"Isaac Asimov",
		"1951",
		"fiction",
		"read",
	))
	assert.Nil(t, err)
}

func TestSelectAllBooks(t *testing.T) {
	expected := []*models.Book{
		models.NewBook(
			"1,000 Year Plan",
			"Isaac Asimov",
			"1951",
			"fiction",
			"read",
		),
		models.NewBook(
			"I Robot",
			"Isaac Asimov",
			"1963",
			"fiction",
			"read",
		),
		models.NewBook(
			"The Collapsing Universe",
			"Isaac Asimov",
			"1977",
			"non-fiction",
			"read",
		),
		models.NewBook(
			"Artificial Intelligence: A Guide for Thinking Humans",
			"Melanie Mitchell",
			"2020",
			"non-fiction",
			"read",
		),
	}

	br := NewBooksRepo(database.Open(bookSchema))

	for _, b := range expected[1:] {
		err := br.Insert(b)
		assert.Nil(t, err)
	}

	actual, err := br.SelectAll()
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)
}

func TestRetrieveBook(t *testing.T) {
	br := NewBooksRepo(database.Open(bookSchema))

	expected := models.NewBook(
		"The End of Eternity",
		"Isaac Asimov",
		"1955",
		"fiction",
		"read",
	)

	err := br.Insert(expected)
	assert.Nil(t, err)

	actual, err := br.Select(expected.ID)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, actual.ID)
}

func TestUpdateBook(t *testing.T) {
	br := NewBooksRepo(database.Open(bookSchema))

	book := models.NewBook(
		"The Time Machine",
		"H. G. Wells",
		"1979",
		"fiction",
		"reading",
	)

	err := br.Insert(book)
	assert.Nil(t, err)

	updated := models.NewBook(
		"The Time Machine",
		"H. G. Wells",
		"1979",
		"fiction",
		"read",
	)

	id, err := br.Update(book.ID, updated)
	assert.Nil(t, err)
	assert.Equal(t, updated.ID, id)
}

func TestDeleteBook(t *testing.T) {
	br := NewBooksRepo(database.Open(bookSchema))

	book := models.NewBook(
		"Chaos: Making a New Science",
		"James Gleick",
		"1988",
		"non-fiction",
		"reading",
	)

	err := br.Insert(book)
	assert.Nil(t, err)

	err = br.Delete(book.ID)
	assert.Nil(t, err)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
