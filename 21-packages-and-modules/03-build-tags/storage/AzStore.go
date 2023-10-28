//go:build az

package storage

import (
	"fmt"
	"os"
)

type AzStore struct {
	ConnString string
}

func New() Store {
	connstring := os.Getenv("AZURE_CONNECTION_STRING")
	if connstring == "" {
		panic("AZURE_CONNECTION_STRING not defined.")
	}
	return &AzStore{
		ConnString: connstring,
	}
}

func (az *AzStore) Upload() {
	fmt.Printf("[%s] Uploading to Azure...\n", az.ConnString)
}

func (az *AzStore) Download() {
	fmt.Printf("[%s] Downloading from Azure...\n", az.ConnString)
}

func (az *AzStore) AzSpecificMethod() {
	fmt.Println("This is Az specific method.")
}
