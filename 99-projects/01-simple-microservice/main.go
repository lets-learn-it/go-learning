package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"avabodha.in/SimpleMicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// register handler
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)

	// http.ListenAndServe(":8080", sm)

	// Instead of ListenAndServe
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until no message in channel
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
