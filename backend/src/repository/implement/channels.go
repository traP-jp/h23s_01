package implement

import (
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23s_01/backend/src/domain"
)

type Channels struct {
	db *sqlx.DB
}

func NewChannels(db *sqlx.DB) *Channels {
	return &Channels{
		db: db,
	}
}

func (c *Channels) RemakeChannelsTable(channels []domain.Channel) error {
	c.db.MustExec("DELETE FROM channels")
	_, err := c.db.NamedExec("INSERT INTO channels (id, name) VALUES (:id, :name)", channels)
	if err != nil {
		return err
	}
	return nil
}

func (c *Channels) GetRandomChannel() (*domain.Channel, error) {
	var ch domain.Channel
	if err := c.db.Get(&ch, "SELECT * FROM channels ORDER BY RAND() LIMIT 1"); err != nil {
		return nil, err
	}
	return &ch, nil
}
