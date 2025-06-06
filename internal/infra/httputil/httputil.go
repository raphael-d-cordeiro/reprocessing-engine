package httputil

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/raphael-d-cordeiro/reprocessing-engine/internal/domain/protocol/httpclient"
)

type Client struct {
	BaseURL   string
	transport http.RoundTripper
	client    *http.Client
}

func NewClient(baseURL string) *Client {
	transport := &http.Transport{
		MaxIdleConns:          500,
		MaxIdleConnsPerHost:   300,
		MaxConnsPerHost:       500,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // For development purposes only
		},
	}
	return &Client{
		BaseURL:   baseURL,
		transport: transport,
		client:    &http.Client{Transport: transport},
	}
}

func (c *Client) MakeRequest(ctx context.Context, params httpclient.RequestParams) (*httpclient.ResponseData, error) {
	// Monta a URL com query params
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = params.URL
	q := u.Query()
	for k, v := range params.QueryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	// Cria o request
	req, err := http.NewRequestWithContext(ctx, params.Method, u.String(), params.Body)
	if err != nil {
		return nil, err
	}
	for k, v := range params.Headers {
		req.Header.Set(k, v)
	}

	// Timeout customizado
	client := c.client
	if params.Timeout > 0 {
		client = &http.Client{Transport: c.transport, Timeout: time.Duration(params.Timeout) * time.Second}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &httpclient.ResponseData{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}
