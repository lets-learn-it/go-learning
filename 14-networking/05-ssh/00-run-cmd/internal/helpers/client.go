package helpers

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type ClientConfig struct {
	Username string
	Key      ssh.Signer
	Host     string
	Port     string
}

func NewSshClient(c ClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(c.Key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := net.JoinHostPort(c.Host, c.Port)
	return ssh.Dial("tcp", address, config)
}
