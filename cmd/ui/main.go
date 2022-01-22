package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func main() {
	tpl := template.Must(template.ParseGlob("templates/*"))

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		br, err := getBooks()
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			err = tpl.ExecuteTemplate(w, "index.html", br)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		br, err := getBookById(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			err = tpl.ExecuteTemplate(w, "book.html", br)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
	})

	addr := "0.0.0.0:9042"
	log.Printf("Serving at %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
