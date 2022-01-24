package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
)

var (
	tplDir = "templates/*"
	tpl *template.Template
)

func main() {
	tpl = template.Must(template.ParseGlob(tplDir))
	if tpl == nil {
		log.Fatalf("template.Must(template.ParseGlob(\"%s\")", tplDir)
	}

	HOST := os.Getenv("HOST")
	if HOST == "" {
		HOST = "0.0.0.0"
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9042"
	}

	addr := fmt.Sprintf("%s:%s", HOST, PORT)

	r := chi.NewRouter()

	r.Post("/", createBook)
	r.Get("/", getBooks)
	r.Get("/{id}", getBookById)
	r.Get("/ping", getHeartbeat)
	r.Get("/create", createBook)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		root := http.Dir("static")
		fsrvr := http.StripPrefix("/static/", http.FileServer(root))
		fsrvr.ServeHTTP(w, r)
	})

	log.Printf("Serving at %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
