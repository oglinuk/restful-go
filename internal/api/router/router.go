package router

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oglinuk/restful-go/internal/api/resources"
)

func NewRouter(db *sql.DB) *chi.Mux {
	env := resources.NewEnv(db)
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60*time.Second))

	r.Get("/", env.Heartbeat)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", env.BookList)
		r.Get("/{id}", env.BookById)
	})

	return r
}
