package router

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	r.Get("/", env.Heartbeat)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", env.BookList)
		r.Get("/{id}", env.BookById)
		r.Delete("/{id}", env.BookById)
	})

	return r
}
