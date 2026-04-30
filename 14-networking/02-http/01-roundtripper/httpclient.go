package main

import (
	"net/http"
	"time"
)

type Option func(*options)

type options struct {
	timeout   time.Duration
	transport http.RoundTripper
}

func WithTimeout(d time.Duration) Option {
	return func(o *options) {
		o.timeout = d
	}
}

func WithTransport(t http.RoundTripper) Option {
	return func(o *options) {
		o.transport = t
	}
}

func New(opts ...Option) *http.Client {
	cfg := &options{
		timeout:   30 * time.Second,
		transport: http.DefaultTransport,
	}
	for _, o := range opts {
		o(cfg)
	}

	transport := cfg.transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}
}
