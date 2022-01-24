package main

import (
	"crypto/tls"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	localIP = "http://0.0.0.0:9001"
	dockerIP = "http://api:9001"
	currentIP = ""

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 15,
	}
)

func init() {
	_, err := http.Get(dockerIP)
	if err != nil {
		log.Printf("%s is not running, defaulting to local ...\n", dockerIP)
		currentIP = localIP
	} else {
		log.Printf("%s is running ...\n", dockerIP)
		currentIP = dockerIP
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	// HTML forms don't allow PUT (or DELETE), so to work around this, we
	// check for a hidden input ("_method") from the form. Need to review.
	// https://www.w3.org/Bugs/Public/show_bug.cgi?id=10671 status is
	// `RESOLVED WONTFIX`?
	isUpdate := r.PostFormValue("_method")
	if isUpdate == "PUT" {
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

		br := &BookResp{}

		err = decodeJSON(br, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("[BookResp]: %v\n", br)

		http.Redirect(w, r, "/"+br.ID, http.StatusMovedPermanently)
		return
	}

	http.Error(
		w,
		fmt.Sprintf("Unsupported METHOD: %s", r.Method),
		http.StatusNotImplemented,
	)
}

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	br := &BookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("[BookResp]: %v\n", br)

	http.Redirect(w, r, "/"+br.ID, http.StatusMovedPermanently)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := http.Get(fmt.Sprintf("%s/books/%s", currentIP, id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	br := &BookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("[BookResp]: %v\n", br)

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

	br := &BooksResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[BooksResp]: %v\n", br)

	err = tpl.ExecuteTemplate(w, "index.html", br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getHeartbeat(w http.ResponseWriter, r *http.Request) {
	hb := &HeartbeatResp{}

	resp, err := http.Get(currentIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = decodeJSON(hb, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[HeartbeatResp]: %v\n", hb)

	fmt.Fprintf(w, fmt.Sprintf("%v", hb))
}

