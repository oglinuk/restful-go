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

func TestInsert(t *testing.T) {
	br := NewBooksRepo(database.Open(defaultBookSchema))
	err := br.Insert(models.NewBook("1,000 Year Plan", "Isaac Asimov", "1951"))
	assert.Nil(t, err)
}

func TestSelectAll(t *testing.T) {
	expected := []*models.Book{
		models.NewBook(
			"1,000 Year Plan",
			"Isaac Asimov",
			"1951",
		),
		models.NewBook(
			"I, Robot",
			"Isaac Asimov",
			"1963",
		),
	}

	br := NewBooksRepo(database.Open(defaultBookSchema))

	for _, b := range expected[1:] {
		br.Insert(b)
	}
	actual, err := br.SelectAll()
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, actual)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
