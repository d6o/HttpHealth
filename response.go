package main

import (
	"encoding/json"
	"net/http"
)

// Response contains all data sent in the response.
type Response struct {
	Pings  map[string]PingResponse
	Writer http.ResponseWriter `json:"-"`
}

// NewResponse initialize the response saving the writer.
func NewResponse(writer http.ResponseWriter) *Response {
	return &Response{Writer: writer}
}

// Write the response to the user.
func (r *Response) Write(status int, pings map[string]PingResponse) {
	r.Writer.WriteHeader(status)
	r.Pings = pings
	res, _ := json.Marshal(&r)
	r.Writer.Write(res)
}
