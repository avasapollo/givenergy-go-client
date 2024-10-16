package inverter

import (
	"net/http"
	"time"
)

type Option func(*options)

type options struct {
	baseURL    string
	httpClient *http.Client
}

func defaultOptions() *options {
	return &options{
		baseURL: "https://api.givenergy.cloud/v1",
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func WithHTTPClient(cl *http.Client) Option {
	return func(o *options) {
		o.httpClient = cl
	}
}

func WithBaseURL(baseURL string) Option {
	return func(o *options) {
		o.baseURL = baseURL
	}
}
