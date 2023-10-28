//go:build aws

package storage

import (
	"fmt"
	"os"
)

type AWSStore struct {
	Id     string
	Secret string
}

func New() *AWSStore {
	id := os.Getenv("AWS_ID")
	if id == "" {
		panic("AWS_ID not found.")
	}

	secret := os.Getenv("AWS_SECRET")
	if id == "" {
		panic("AWS_SECRET not found.")
	}

	return &AWSStore{
		Id:     id,
		Secret: secret,
	}
}

func (aws *AWSStore) Upload() {
	fmt.Printf("[%s] Uploading to AWS...\n", aws.Id)
}

func (aws *AWSStore) Download() {
	fmt.Printf("[%s] Downloading from AWS...\n", aws.Id)
}
