package helpers

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func LoadSshSigner(keyPath string, passphrase string) (ssh.Signer, error) {
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("Failed to read private key: %s", err)
	}

	var key ssh.Signer

	if passphrase == "" {
		key, err = ssh.ParsePrivateKey(keyBytes)
	} else {
		key, err = ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte(passphrase))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	return key, nil
}
