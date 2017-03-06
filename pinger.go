package main

import (
	"net/http"
	"time"
)

type Pinger interface {
	Ping() (int, error)
	URL() string
}

type HeadPing struct {
	url string
}

type PingResponse struct {
	Code    int
	Message string
}

func NewHeadPing(url string) *HeadPing {
	return &HeadPing{url: url}
}

func (h *HeadPing) Ping() (PingResponse, error) {
	p := PingResponse{
		Code:    0,
		Message: "UP",
	}

	client := http.Client{Timeout: 5 * time.Second}
	r, err := client.Head(h.url)
	if err != nil {
		p.Message = err.Error()
	} else {
		defer r.Body.Close()
		p.Code = r.StatusCode
	}

	return p, nil
}

func (h *HeadPing) URL() string {
	return h.url
}
