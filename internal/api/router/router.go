package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/oglinuk/restful-go/internal/api/resources"
)

func New(db *sql.DB) *mux.Router {
	env := resources.NewEnv(db)
	router := mux.NewRouter()

	router.HandleFunc("/", env.Heartbeat)
	router.HandleFunc("/books", env.BookList)

	return router
}
