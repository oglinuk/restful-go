package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := http.Get(fmt.Sprintf("%s/books/%s", currentIP, id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	br := &bookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[getBookById::bookResp]: %v\n", br)

	err = tpl.ExecuteTemplate(w, "book.html", br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("%s/books", currentIP))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	br := &booksResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[getBooks::booksResp]: %v\n", br)

	err = tpl.ExecuteTemplate(w, "index.html", br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
