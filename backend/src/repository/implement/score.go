package implement

import (
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23s_01/backend/src/domain"
)

type Score struct {
	db *sqlx.DB
}

func NewScore(db *sqlx.DB) *Score {
	return &Score{
		db: db,
	}
}

func (s *Score) AddScore(score *domain.Score) error {
	_, err := s.db.NamedExec("INSERT INTO scores (id, user_id, score) VALUES (:id, :user_id, :score)", score)
	if err != nil {
		return err
	}
	return nil
}
