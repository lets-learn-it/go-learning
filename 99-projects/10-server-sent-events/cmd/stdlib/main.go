package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Line struct {
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Println("WRITER")
	var p1, p2 Point
	var line Line

	for {
		select {
		case <-r.Context().Done():
			log.Println("Client closed connection")
			return
		default:
			p1.X = rand.Intn(400)
			p1.Y = rand.Intn(400)

			p2.X = rand.Intn(400)
			p2.Y = rand.Intn(400)

			line.P1 = p1
			line.P2 = p2

			data, err := json.Marshal(line)
			if err != nil {
				log.Println(err)
			}

			fmt.Fprintf(w, "data: %s\n\n", data)

			flusher.Flush()
			time.Sleep(time.Second)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/api/sse", eventsHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
