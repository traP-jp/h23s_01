package handler

import (
	"fmt"
	"net/http"
	"strconv"
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

type result struct {
	UserName  string    `json:"user_name"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func (sh *ScoreHandler) getScoreRankingHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	strLimit := c.QueryParam("limit")
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		if strLimit == "" {
			limit = 10
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}
	if limit < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "limit should be positive number")
	}

	ranking, err := sh.sr.GetScoreRandking(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	var responseRanking []result

	for _, r := range ranking {
		userInfo, err := sh.tc.GetUserInfo(token, r.UserId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
		}
		responseRanking = append(responseRanking, result{UserName: userInfo.Name, Score: r.Score, CreatedAt: r.CreatedAt})
	}

	return c.JSON(http.StatusOK, responseRanking)
}
