package repository

import (
	"sync"

	"library-api/pkg/library"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	repo library.Repository
	once sync.Once
)

type repository struct {
	conn *mongo.Collection
}

func Must(mongoConnection *mongo.Collection) library.Repository {
	once.Do(func() {
		repo = &repository{
			conn: mongoConnection,
		}
	})

	return repo
}
