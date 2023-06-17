package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/repository"
	"github.com/traP-jp/h23s_01/backend/src/traq"
)

type channelHandler struct {
	tc traq.TraqClient
	cr repository.ChannelsRepository
}

func NewChannelHandler(tc traq.TraqClient, cr repository.ChannelsRepository) *channelHandler {
	return &channelHandler{
		tc: tc,
		cr: cr,
	}
}

func (ch *channelHandler) patchChennelsHandler(c echo.Context) error {
	root := c.QueryParam("root")
	if root == "" {
		root = "gps+times"
	}
	root = strings.ReplaceAll(root, "+", "/")

	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	channels, err := ch.tc.GetAllChannels(token, root)
	if errors.Is(err, traq.ErrNoChannel) {
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	if err := ch.cr.RemakeChannelsTable(channels); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %v", err))
	}

	return c.String(http.StatusOK, "ok")
}
