package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pskp-95/ssh-run-cmd/internal/helpers"
)

func main() {
	keyPath := flag.String("key", "/Users/paripati/.ssh/jumphost", "private key location")
	passphrase := flag.String("passphrase", "", "passphrase for key")
	host := flag.String("host", "127.0.0.1", "host ip or dns")
	port := flag.String("port", "22", "port on which ssh server running")
	command := flag.String("command", "ls", "command to run on server")
	username := flag.String("username", "root", "user")

	flag.Parse()

	key, err := helpers.LoadSshSigner(*keyPath, *passphrase)
	if err != nil {
		log.Fatalf("failed to load key %s: %s", key, err)
	}

	client, err := helpers.NewSshClient(helpers.ClientConfig{
		Username: *username,
		Host:     *host,
		Port:     *port,
		Key:      key,
	})
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}

	fmt.Printf("Connected to %s\n", host)
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %s", err)
	}
	defer session.Close()

	// Run the command on the remote machine
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get stdout: %s", err)
	}
	if err := session.Start(*command); err != nil {
		log.Fatalf("failed to start command: %s", err)
	}

	_, err = io.Copy(os.Stdout, stdout)
	if err != nil {
		log.Fatalf("failed to read stdout: %s", err)
	}
	if err := session.Wait(); err != nil {
		log.Fatalf("command failed: %s", err)
	}
}
