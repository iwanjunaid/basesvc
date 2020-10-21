package model

import (
	"time"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    Author
}
