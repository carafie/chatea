package web

import (
	"encoding/json/v2"
	"net/http"
)

type response struct {
	statusCode int
	body       any
}

func (r response) respond(w http.ResponseWriter) {
	if r.body != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(r.statusCode)
		_ = json.MarshalWrite(w, r.body)
	} else {
		w.WriteHeader(r.statusCode)
	}
}

type errorBody struct {
	Code string `json:"code"`
}
