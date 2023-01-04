package library

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Book struct {
	ID          uuid.UUID
	Name        string
	Description string
	Writer      string
	ReleaseDate time.Time
	CreatedAt   time.Time
	DelectedAt  time.Time
}
