package repository

import "github.com/traP-jp/h23s_01/backend/src/domain"

type ScoreRepository interface {
	AddScore(score *domain.Score) error
}
