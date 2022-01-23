package main

import (
	"crypto/tls"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func createBook(r *http.Request) (string, error) {
	payload := map[string]string{
		"title": r.PostFormValue("title"),
		"author": r.PostFormValue("author"),
		"published": r.PostFormValue("published"),
		"genre": r.PostFormValue("genre"),
		"readstatus": r.PostFormValue("readstatus"),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/books", currentIP),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	br := &BookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("[BookResp]: %v\n", br)

	return br.ID, nil
}

func updateBookById(r *http.Request) (*BookResp, error) {
	payload := map[string]string{
		"title": r.PostFormValue("title"),
		"author": r.PostFormValue("author"),
		"published": r.PostFormValue("published"),
		"genre": r.PostFormValue("genre"),
		"readstatus": r.PostFormValue("readstatus"),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/books/%s", currentIP, r.PostFormValue("id")),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	br := &BookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[BookResp]: %v\n", br)

	return br, nil
}

func getBookById(id string) (*BookResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/books/%s", currentIP, id))
	if err != nil {
		return nil, err
	}

	br := &BookResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[BookResp]: %v\n", br)

	return br, nil
}

func getBooks() (*BooksResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/books", currentIP))
	if err != nil {
		return nil, err
	}

	br := &BooksResp{}

	err = decodeJSON(br, resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[BooksResp]: %v\n", br)

	return br, nil
}

func getHeartbeat() (*HeartbeatResp, error) {
	hb := &HeartbeatResp{}

	resp, err := http.Get(currentIP)
	if err != nil {
		return nil, err
	}

	err = decodeJSON(hb, resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[HeartbeatResp]: %v\n", hb)

	return hb, nil
}

