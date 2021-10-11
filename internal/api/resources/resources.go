package resources

import (
	"database/sql"

	"github.com/oglinuk/restful-go/internal/pkg/repositories"
)

type Env struct {
	Books *repositories.BooksRepo
}

func NewEnv(db *sql.DB) *Env {
	return &Env{
		Books: repositories.NewBooksRepo(db),
	}
}
