package httpclient

import (
	"context"
	"io"
)

// RequestParams representa os parâmetros de uma requisição HTTP genérica.
type RequestParams struct {
	Method      string
	URL         string
	Body        io.Reader
	Headers     map[string]string
	QueryParams map[string]string
	Timeout     int // Timeout in seconds
}

// ResponseData representa a resposta de uma requisição HTTP genérica.
type ResponseData struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
}

// HttpInterface define o contrato para um client HTTP genérico.
type HttpInterface interface {
	MakeRequest(ctx context.Context, data RequestParams) (*ResponseData, error)
}
