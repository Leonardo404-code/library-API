package repository

import (
	"library-api/pkg/library"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	repo library.Repository
	once sync.Once
)

type repository struct {
	conn *mongo.Client
}

func Must(mongoConnection *mongo.Client) library.Repository {
	once.Do(func() {
		repo = &repository{
			conn: mongoConnection,
		}
	})

	return repo
}
