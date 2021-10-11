package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

// TestOpen tests Open using a simple single table/column schema
func TestOpen(t *testing.T) {
	db := Open(`CREATE TABLE tblTest(id INTEGER PRIMARY KEY AUTOINCREMENT);`)
	assert.NotNil(t, db)

	t.Cleanup(func() {
		cfg := config.Get()
		os.Remove(cfg.Name)
		os.Remove(cfg.Database.File)
	})
}
