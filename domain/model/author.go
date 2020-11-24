package model

import (
	uuid "github.com/satori/go.uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt int64
	UpdatedAt int64
}
