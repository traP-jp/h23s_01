package handler

import (
	"net/http"

	"github.com/traP-jp/h23s_01/backend/src/traq"

	"github.com/labstack/echo/v4"
)

type traqClient struct {
	tc traq.TraqClient
}

func NewTraqClient(tc traq.TraqClient) *traqClient {
	return &traqClient{
		tc: tc,
	}
}

func (tc *traqClient) getMeHandler(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	user, err := tc.tc.GetMe(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
