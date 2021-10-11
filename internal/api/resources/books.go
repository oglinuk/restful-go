package resources

import (
	"net/http"
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
