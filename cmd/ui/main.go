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

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// This is dumb, but required since PUT is not allowed in HTML forms
		isUpdate := r.PostFormValue("_method")

		if isUpdate != "" {
			b, err := updateBookById(r)
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				http.Redirect(w, r, "/"+b.ID, http.StatusMovedPermanently)
			}
			return
		}

		id, err := createBook(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Redirect(w, r, "/"+id, http.StatusMovedPermanently)
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

	r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		root := http.Dir("static")
		fsrvr := http.StripPrefix("/static/", http.FileServer(root))
		fsrvr.ServeHTTP(w, r)
	})

	addr := "0.0.0.0:9042"
	log.Printf("Serving at %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
