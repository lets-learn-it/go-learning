package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// Header round tripper
// It will add headers per request.
type headerRoundTripper struct {
	next    http.RoundTripper
	headers map[string]string
}

func NewHeaderRoundTripper(next http.RoundTripper, headers map[string]string) *headerRoundTripper {
	return &headerRoundTripper{
		next:    next,
		headers: headers,
	}
}

func (h *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("[headerRoundTripper] before request sent")

	for k, v := range h.headers {
		req.Header.Add(k, v)
	}

	resp, err := h.next.RoundTrip(req)

	fmt.Println("[headerRoundTripper] after request sent")

	return resp, err
}

// Basic Auth round tripper
// It will add Authorization header per request.
type basicAuthRoundTripper struct {
	next     http.RoundTripper
	username string
	password string
}

func NewBasicAuthRoundTripper(next http.RoundTripper, username, password string) *basicAuthRoundTripper {
	return &basicAuthRoundTripper{
		next:     next,
		username: username,
		password: password,
	}
}

func (h *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("[basicAuthRoundTripper] before request sent")

	header := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", h.username, h.password)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", header))

	resp, err := h.next.RoundTrip(req)

	fmt.Println("[basicAuthRoundTripper] after request sent")

	return resp, err
}
