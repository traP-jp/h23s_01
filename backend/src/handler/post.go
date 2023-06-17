package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type scoreBinder struct {
	Score int `json:"score"`
}

func (tc *traqClient) postScoreHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	var score scoreBinder
	c.Bind(&score)
	if err := tc.tc.PostScore(token, score.Score); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	return c.String(http.StatusOK, "ok")
}
