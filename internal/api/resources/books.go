package resources

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// BookList gets and returns all available books in the database
func (env *Env) BookList(w http.ResponseWriter, r *http.Request) {
	books, err := env.Books.SelectAll()
	if err != nil {
		JSONIFY(w, http.StatusInternalServerError, JSON{"error": err.Error()})
	} else {
		JSONIFY(w, http.StatusOK, JSON{"books": books})
	}
}

// BookById gets and returns the book with the given {id}
func (env *Env) BookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		JSONIFY(
			w,
			http.StatusBadRequest,
			JSON{"error": "id was not given"},
		)
		return
	}

	book, err := env.Books.Select(id)
	if err != nil {
		JSONIFY(w, http.StatusInternalServerError, JSON{"error": err.Error()})
	} else {
		JSONIFY(w, http.StatusOK, JSON{
			"id": book.ID,
			"title": book.Title,
			"author": book.Author,
			"published": book.Published,
			"genre": book.Genre,
			"readstatus": book.ReadStatus,
		})
	}
}
