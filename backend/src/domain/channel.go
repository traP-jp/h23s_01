package domain

import "github.com/google/uuid"

type Channel struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}
