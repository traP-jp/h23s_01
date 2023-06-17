package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	User      string
	MessageId uuid.UUID
	Channel   string
	Content   string
	CreatedAt time.Time
	Ika       bool
	Shika     bool
	Meka      bool
}
