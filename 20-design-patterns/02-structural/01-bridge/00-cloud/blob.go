package main

import "fmt"

type AzureStorage struct {
	storageAccount string
	container      string
}

func (b *AzureStorage) Upload() {
	fmt.Printf("uploading to storage account %s & container %s\n", b.storageAccount, b.container)
}
