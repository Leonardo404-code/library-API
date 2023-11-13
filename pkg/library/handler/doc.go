package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type (
	BookResponseDoc struct {
		ID          uuid.UUID `example:"36e4379a-be55-462d-ab9e-da0c532623ea"`
		Title       string    `example:"Test title"                           json:"title"`
		Writer      string    `example:"Jhonn Doe"                            json:"write"`
		ReleaseDate string    `example:"2023"                                 json:"release_date"`
		Description string    `example:"Some description here"                json:"description"`
		Genre       string    `example:"notice"                               json:"genre"`
		CreatedAt   time.Time `example:"2006-01-02T15:04:05Z"                 json:"created_at"   format:"date"`
		UpdatedAt   time.Time `example:"2006-01-02T15:04:05Z"                 json:"updated_at"   format:"date"`
		DeletedAt   time.Time `example:"2006-01-02T15:04:05Z"                 json:"deleted_at"   format:"date"`
	}

	BookRequestDoc struct {
		Title       string `json:"title"        example:"Test title"`
		Writer      string `json:"write"        example:"Jhonn Doe"`
		ReleaseDate string `json:"release_date" example:"2023"`
		Description string `json:"description"  example:"Some description here"`
		Genre       string `json:"genre"        example:"notice"`
	}
)
