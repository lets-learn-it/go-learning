package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	transport := NewHeaderRoundTripper(
		NewBasicAuthRoundTripper(http.DefaultTransport, "pskp", "******"),
		map[string]string{
			"MyHeader":   "Yes",
			"YourHeader": "No",
		},
	)

	httpclient := New(
		WithTimeout(5*time.Second),
		WithTransport(transport),
	)

	request, _ := http.NewRequest("GET", "https://httpbin.org/basic-auth/pskp/******", nil)

	resp, err := httpclient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
}
