package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	tpl := template.Must(template.ParseGlob("templates/*"))

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		br, err := getBookById(mux.Vars(r)["id"])
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
	log.Fatal(http.ListenAndServe(addr, router))
}
