package implement

import (
	"fmt"

	"github.com/google/uuid"
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

func (s *Score) RegisterScore(score *domain.Score) error {
	_, err := s.db.NamedExec("INSERT INTO scores (id, user_id, score) VALUES (:id, :user_id, :score)", score)
	if err != nil {
		return err
	}
	return nil
}

func (s *Score) GetHighestScore(userId uuid.UUID) (*domain.Score, error) {
	var highScore domain.Score
	err := s.db.Get(&highScore, "SELECT score, created_at FROM scores WHERE user_id = ? ORDER BY score DESC LIMIT 1", userId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", highScore)
	return &highScore, nil
}

func (s *Score) GetScoreRandking(limit int) ([]domain.Score, error) {
	var ranking []domain.Score
	if err := s.db.Select(&ranking, "SELECT * FROM scores ORDER BY score DESC LIMIT ?", limit); err != nil {
		return nil, err
	}
	return ranking, nil
}
