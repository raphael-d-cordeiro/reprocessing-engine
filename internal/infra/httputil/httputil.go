package httputil

import (
	"net/http"
	"time"
)

type HttpClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}
