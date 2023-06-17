package repository

import "github.com/traP-jp/h23s_01/backend/src/domain"

type ChannelsRepository interface {
	RemakeChannelsTable([]domain.Channel) error
	GetRandomChannel() (*domain.Channel, error)
}
