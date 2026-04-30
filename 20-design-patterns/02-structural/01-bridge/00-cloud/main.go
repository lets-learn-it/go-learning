package main

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
