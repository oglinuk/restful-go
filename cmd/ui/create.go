package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createBook(w http.ResponseWriter, r *http.Request) {
	// HTML forms don't allow PUT (or DELETE), so to work around this, we
	// check for a hidden input ("_method") from the form. Need to review.
	// https://www.w3.org/Bugs/Public/show_bug.cgi?id=10671 status is
	// `RESOLVED WONTFIX`?
	isUpdate := r.PostFormValue("_method")
	if isUpdate == "PUT" || r.Method == "PUT" {
		updateBookById(w, r)
		return
	}

	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method == "POST" {
		payload := map[string]string{
			"title": r.PostFormValue("title"),
			"author": r.PostFormValue("author"),
			"published": r.PostFormValue("published"),
			"genre": r.PostFormValue("genre"),
			"readstatus": r.PostFormValue("readstatus"),
		}

		data, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/books", currentIP),
			bytes.NewBuffer(data),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		req.Header.Set("Content-Type", "application/json; charset=utf-8")

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		br := &bookResp{}

		err = decodeJSON(br, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Printf("[createBooks::bookResp]: %v\n", br)

		http.Redirect(w, r, "/"+br.ID, http.StatusMovedPermanently)
		return
	}

	http.Error(
		w,
		fmt.Sprintf("Unsupported METHOD: %s", r.Method),
		http.StatusNotImplemented,
	)
}
