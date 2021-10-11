package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getBooks() (*BooksResp, error) {
	resp, err := http.Get("http://0.0.0.0:9001/books")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	br := &BooksResp{}

	err = json.NewDecoder(resp.Body).Decode(&br)
	if err != nil {
		return nil, err
	}
	log.Printf("[BooksResp]: %v\n", br)

	return br, nil
}

func getHeartbeat() (*HeartbeatResp, error) {
	hb := &HeartbeatResp{}

	resp, err := http.Get("http://0.0.0.0:9001")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&hb)
	if err != nil {
		return nil, err
	}
	log.Printf("[HeartbeatResp]: %v\n", hb)

	return hb, nil
}

