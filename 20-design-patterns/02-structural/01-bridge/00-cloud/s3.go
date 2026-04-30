package main

import "fmt"

type S3 struct {
	bucket string
}

func (b *S3) Upload() {
	fmt.Printf("uploading to s3 %s\n", b.bucket)
}
