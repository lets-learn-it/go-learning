package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s\n", string(req))
	io.WriteString(w, string(req))
}

func main() {
	// get port
	sPort := os.Getenv("ECHO_PORT")
	if sPort == "" {
		sPort = "8080"
	}

	port, err := strconv.Atoi(sPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to convert ECHO_PORT to int. %s\n", sPort)
		os.Exit(1)
	}

	if port <= 1024 || port > 65535 {
		fmt.Fprintf(os.Stderr, "ECHO_PORT must be between 1025 - 65535.\n")
		os.Exit(1)
	}

	http.HandleFunc("/", getRoot)

	err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "./ca.crt", "./ca.key", nil)
	if err != nil {
		panic(err)
	}
}
