package util

import (
	"net/http"
	"time"
)

func NewDefaultHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    100,
			MaxConnsPerHost: 1000,
			IdleConnTimeout: 10 * time.Second,
		},
	}
}
