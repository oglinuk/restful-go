package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func updateBookById(w http.ResponseWriter, r *http.Request) {
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
		"PUT",
		fmt.Sprintf("%s/books/%s", currentIP, r.PostFormValue("id")),
		bytes.NewBuffer(data),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil || resp == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	br := &bookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[updateBookById::bookResp]: %v\n", br)

	http.Redirect(w, r, "/"+br.ID, http.StatusMovedPermanently)
}
