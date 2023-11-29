package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type (
	BookResponseDoc struct {
		ID          uuid.UUID `example:"36e4379a-be55-462d-ab9e-da0c532623ea"`
		Title       string    `json:"title" example:"Test title"`
		Writer      string    `json:"write" example:"Jhonn Doe"`
		ReleaseDate string    `json:"release_date" example:"2023"`
		Description string    `json:"description" example:"Some description here"`
		Genre       string    `json:"genre" example:"notice"`
		CreatedAt   time.Time `json:"created_at"   format:"date" example:"2006-01-02T15:04:05Z"`
		UpdatedAt   time.Time `json:"updated_at"   format:"date" example:"2006-01-02T15:04:05Z"`
		DeletedAt   time.Time `json:"deleted_at"   format:"date" example:"2006-01-02T15:04:05Z"`
	}

	BookRequestDoc struct {
		BookInfo bookInfo `formData:"json" bidding:"required"`
	}

	bookInfo struct {
		Title       string `json:"title"        example:"Test title"`
		Writer      string `json:"write"        example:"Jhonn Doe"`
		ReleaseDate string `json:"release_date" example:"2023"`
		Description string `json:"description"  example:"Some description here"`
		Genre       string `json:"genre"        example:"notice"`
	}
)
