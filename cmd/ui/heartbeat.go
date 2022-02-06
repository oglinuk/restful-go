package main

import (
	"log"
	"net/http"
)

func getHeartbeat(w http.ResponseWriter, r *http.Request) {
	hb := &heartbeatResp{}

	resp, err := http.Get(currentIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = decodeJSON(hb, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[HeartbeatResp]: %v\n", hb)

	err = tpl.ExecuteTemplate(w, "ping.html", hb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

