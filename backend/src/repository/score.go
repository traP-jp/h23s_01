package repository

import (
	"github.com/google/uuid"
	"github.com/traP-jp/h23s_01/backend/src/domain"
)

type ScoreRepository interface {
	RegisterScore(score *domain.Score) error
	GetHighestScore(userId uuid.UUID) (*domain.Score, error)
	GetScoreRandking(limit int) ([]domain.Score, error)
}
