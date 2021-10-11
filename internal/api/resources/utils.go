package resources

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type JSON map[string]interface{}

func JSONIFY(w http.ResponseWriter, code int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func Record(r *http.Request, handler http.HandlerFunc) *http.Response {
	rr := httptest.NewRecorder()
	handler(rr, r)
	return rr.Result()
}
