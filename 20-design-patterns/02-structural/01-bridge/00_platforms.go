package main

import "fmt"

type IStorage interface {
	Upload()
}

// first concrete implementation
type S3 struct {
	bucket string
}

func (b *S3) Upload() {
	fmt.Printf("uploading to s3 %s\n", b.bucket)
}

// second concrete implementation
type AzureStorage struct {
	storageAccount string
	container      string
}

func (b *AzureStorage) Upload() {
	fmt.Printf("uploading to storage account %s & container %s\n", b.storageAccount, b.container)
}

// my application
type MyApp struct {
	storage IStorage
}

func (m *MyApp) uploadUserData() {
	m.storage.Upload()
}

func main() {
	s3 := &S3{bucket: "my_bucket"}
	azStorage := &AzureStorage{storageAccount: "my_account", container: "some_container"}

	myApp1 := &MyApp{storage: s3}
	myApp2 := &MyApp{storage: azStorage}

	myApp1.uploadUserData()
	myApp2.uploadUserData()
}
