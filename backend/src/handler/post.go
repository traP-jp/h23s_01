package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/config"
	"github.com/traP-jp/h23s_01/backend/src/repository"
)

type scoreRequestBody struct {
	Score int    `json:"score"`
	Id    string `json:"id"`
}

type postHandler struct {
	ur repository.UsersRepository
}

func NewPostHandler(ur repository.UsersRepository) *postHandler {
	return &postHandler{
		ur: ur,
	}
}

func (ph *postHandler) postScoreHandler(c echo.Context) error {
	var req scoreRequestBody
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if _, err := uuid.Parse(req.Id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to parse user uuid")
	}

	name, err := ph.ur.GetUserNameById(req.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = postWebhook(fmt.Sprintf(message, name, name, req.Score))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "ok")
}

var writer = config.GetWebhookConfig()

const message = `:@%s: @%s は「[いかしかめかアクティビティズ](https://activities-traq.vercel.app/)」で %d 点獲得しました！`

func postWebhook(content string) error {
	if writer == nil {
		fmt.Println(content)
		return nil
	}
	writer.SetEmbed(true)
	_, err := fmt.Fprint(writer, content)
	return err
}
