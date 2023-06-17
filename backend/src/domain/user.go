package domain

import "github.com/google/uuid"

type User struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}
