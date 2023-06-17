package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	User      string    `json:"user,omitempty" db:"user"`
	MessageId uuid.UUID `json:"messageId,omitempty" db:"message_id"`
	Channel   string    `json:"channel,omitempty" db:"channel"`
	Content   string    `json:"content,omitempty" db:"content"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
	Ika       bool      `json:"ika,omitempty" db:"ika"`
	Shika     bool      `json:"shika,omitempty" db:"shika"`
	Meka      bool      `json:"meka,omitempty" db:"meka"`
}
