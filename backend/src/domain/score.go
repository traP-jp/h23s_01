package domain

import (
	"time"

	"github.com/google/uuid"
)

type Score struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	Score     int       `json:"score" db:"score"`
	CreatedAt time.Time `db:"created_at"`
}
