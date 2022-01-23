package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
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

func ChiURLParams(kvs map[string]string, r *http.Request) *http.Request {
	rctx := chi.NewRouteContext()

	for k := range kvs {
		rctx.URLParams.Add(k, kvs[k])
	}

	r = r.WithContext(context.WithValue(
			r.Context(),
			chi.RouteCtxKey,
			rctx,
	))

	return r
}
