package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/config"
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
	err = c.Bind(&score)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := tc.tc.GetMe(token)
	name := user.Name
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = postWebhook(fmt.Sprintf(message, name, name, score.Score))

	return c.String(http.StatusOK, "ok")
}

var writer = config.GetWebhookConfig()

const message = `:@%s: @%s は「[いかしかめかアクティビティズ](https://activities-ecru.vercel.app/)」で %d 点獲得しました！`

func postWebhook(content string) error {
	if writer == nil {
		fmt.Println(content)
		return nil
	}
	writer.SetEmbed(true)
	_, err := fmt.Fprint(writer, content)
	return err
}
