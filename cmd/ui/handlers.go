package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	localIP = "http://0.0.0.0:9001"
	dockerIP = "http://api:9001"
	currentIP = ""
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

func getBooks() (*BooksResp, error) {
	resp, err := http.Get(fmt.Sprintf("%s/books", currentIP))
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

	resp, err := http.Get(currentIP)
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

