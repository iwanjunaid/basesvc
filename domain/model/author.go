package model

import (
	"time"
)

type Author struct {
	ID        uint   `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:email`
	CreatedAt time.Time
	UpdatedAt time.Time
}
