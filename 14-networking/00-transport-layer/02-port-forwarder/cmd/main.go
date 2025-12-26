package main

import (
	"flag"
	"log"

	"github.com/PSKP-95/port-forwarder/internal/tunnel"
)

func isPortValid(port int) bool {
	if port < 0 || port > 65535 {
		return false
	}

	return true
}

func main() {
	var sourcePort, destPort int
	flag.IntVar(&sourcePort, "s", 8080, "An source port")
	flag.IntVar(&destPort, "d", 8081, "An destination port")

	flag.Parse()

	if !isPortValid(sourcePort) || !isPortValid(destPort) {
		log.Fatalf("source port (%d) or destination port (%d) not valid", sourcePort, destPort)
	}

	log.Printf("forwarding %d -> %d", sourcePort, destPort)

	t := tunnel.New(sourcePort, destPort)

	if err := t.Init(); err != nil {
		log.Fatalf("failed to initialize tunnel: %v", err)
	}

	t.Start()
}
