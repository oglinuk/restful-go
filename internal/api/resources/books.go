package resources

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oglinuk/restful-go/internal/pkg/models"
)

// CreateBook
func (env *Env) CreateBook(w http.ResponseWriter, r *http.Request) {
	b := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		JSONIFY(
			w,
			http.StatusBadRequest,
			JSON{"error": err.Error()},
		)
		return
	}

	book := models.NewBook(
		b.Title,
		b.Author,
		b.Published,
		b.Genre,
		b.ReadStatus,
	)

	log.Printf("[REST] Creating %v\n", book)

	err = env.Books.Insert(book)
	if err != nil {
		JSONIFY(
			w,
			http.StatusInternalServerError,
			JSON{"error": err.Error()},
		)
		return
	}

	JSONIFY(w, http.StatusOK, JSON{"id": book.ID})
}

// BookList gets and returns all available books in the database
func (env *Env) BookList(w http.ResponseWriter, r *http.Request) {
	books, err := env.Books.SelectAll()
	if err != nil {
		JSONIFY(w, http.StatusInternalServerError, JSON{"error": err.Error()})
	} else {
		JSONIFY(w, http.StatusOK, JSON{"books": books})
	}
}

// BookById gets and returns the book with the given {id} (if a GET),
// updates and returns the book id (if a PUT), or deletes the book (if a
// DELETE)
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

	if r.Method == "GET" {
		log.Printf("[REST] Getting %s\n", id)

		book, err := env.Books.Select(id)
		if err != nil {
			JSONIFY(
				w,
				http.StatusInternalServerError,
				JSON{"error": err.Error()},
			)
			return
		}

		JSONIFY(w, http.StatusOK, JSON{
			"id": book.ID,
			"title": book.Title,
			"author": book.Author,
			"published": book.Published,
			"genre": book.Genre,
			"readstatus": book.ReadStatus,
		})
		return
	}

	if r.Method == "PUT" {
		b := &models.Book{}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			JSONIFY(
				w,
				http.StatusBadRequest,
				JSON{"error": err.Error()},
			)
			return
		}

		book := models.NewBook(
			b.Title,
			b.Author,
			b.Published,
			b.Genre,
			b.ReadStatus,
		)

		log.Printf("[REST] Updating [%s]: %v\n", id, book)

		id, err := env.Books.Update(id, book)
		if err != nil {
			JSONIFY(
				w,
				http.StatusInternalServerError,
				JSON{"error": err.Error()},
			)
			return
		}

		JSONIFY(w, http.StatusOK, JSON{"id": id})
	}

	if r.Method == "DELETE" {
		log.Printf("[REST] Deleting %s\n", id)

		err := env.Books.Delete(id)
		if err != nil {
			JSONIFY(
				w,
				http.StatusInternalServerError,
				JSON{"error": err.Error()},
			)
			return
		}

		JSONIFY(w, http.StatusOK, JSON{"status": "success"})
	}
}
