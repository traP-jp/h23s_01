package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/domain"
	"github.com/traP-jp/h23s_01/backend/src/repository"
	"github.com/traP-jp/h23s_01/backend/src/traq"
)

type ScoreHandler struct {
	tc traq.TraqClient
	sr repository.ScoreRepository
}

func NewScoreHandler(tc traq.TraqClient, sr repository.ScoreRepository) *ScoreHandler {
	return &ScoreHandler{
		tc: tc,
		sr: sr,
	}
}

func (sh *ScoreHandler) registerScoreHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := sh.tc.GetMe(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	var score domain.Score
	c.Bind(&score)
	score.Id = uuid.New()
	score.UserId = user.Id
	if err := sh.sr.RegisterScore(&score); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]int{"score": score.Score})
}

func (sh *ScoreHandler) getHighestScoreHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := sh.tc.GetMe(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	highScore, err := sh.sr.GetHighestScore(user.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp := struct {
		Score     int       `json:"score"`
		CreatedAt time.Time `json:"created_at"`
	}{
		highScore.Score,
		highScore.CreatedAt,
	}

	return c.JSON(http.StatusOK, resp)
}
