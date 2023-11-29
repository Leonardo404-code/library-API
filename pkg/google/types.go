package google

import "cloud.google.com/go/storage"

type client struct {
	storage    *storage.Client
	bucketName string
	bucketURL  string
}
