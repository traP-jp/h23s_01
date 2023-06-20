package handler

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23s_01/backend/src/domain"
	"github.com/traP-jp/h23s_01/backend/src/reading"
	"github.com/traP-jp/h23s_01/backend/src/repository"
	"github.com/traP-jp/h23s_01/backend/src/traq"
)

type messageHandler struct {
	tc traq.TraqClient
	cr repository.ChannelsRepository
	ur repository.UsersRepository
	r  reading.Reading
}

func NewMessageHandler(tc traq.TraqClient, cr repository.ChannelsRepository, ur repository.UsersRepository, r reading.Reading) *messageHandler {
	return &messageHandler{
		tc: tc,
		cr: cr,
		ur: ur,
		r:  r,
	}
}

type getMessagesRes struct {
	Count    int              `json:"count,omitempty"`
	Messages []domain.Message `json:"messages,omitempty"`
}

func (mh *messageHandler) getMessagesHandler(c echo.Context) error {
	countStr := c.QueryParam("count")
	var count int
	if countStr == "" {
		count = 150
	} else {
		c, err := strconv.Atoi(countStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		count = c
	}

	messagesHit := make([]domain.Message, 0, count)
	messagesNot := make([]domain.Message, 0, count)

	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, key, token)

	result := concurrentTask(ctx, mh)

	for {
		mes := <-result
		if mes.Ika || mes.Shika || mes.Meka {
			messagesHit = append(messagesHit, mes)
		} else {
			messagesNot = append(messagesNot, mes)
		}
		if len(messagesHit)+len(messagesNot) >= count && len(messagesHit) >= count*3/10 && len(messagesHit) <= count*4/10 {
			cancel()
			break
		}
	}

	messagesRes := make([]domain.Message, 0, count)
	messagesRes = append(messagesRes, messagesHit...)
	messagesRes = append(messagesRes, messagesNot...)
	messagesRes = messagesRes[:count]

	return c.JSON(http.StatusOK, getMessagesRes{
		Count:    len(messagesRes),
		Messages: messagesRes,
	})
}

const LENGTH_LIMIT = 50

var (
	key         = tokenKey{}
	ikaRegexp   = regexp.MustCompile(`(い|イ|ｲ)(か|カ|ｶ)`)
	shikaRegexp = regexp.MustCompile(`(し|シ|ｼ)(か|カ|ｶ)`)
	mekaRegexp  = regexp.MustCompile(`(め|メ|ﾒ)(か|カ|ｶ)`)
)

type tokenKey struct{}

func checkIkaShikaMeka(content string, reg *regexp.Regexp) bool {
	return reg.MatchString(content)
}

func checkLength(content string) bool {
	return utf8.RuneCountInString(content) <= LENGTH_LIMIT
}

func concurrentTask(ctx context.Context, mh *messageHandler) <-chan domain.Message {
	result := make(chan domain.Message)
	for i := 0; i < 10; i++ {
		go getMessage(ctx, mh, result)
	}
	return result
}

func getMessage(ctx context.Context, mh *messageHandler, ch chan domain.Message) {
	token := ctx.Value(key).(string)
	for {
		channel, _ := mh.cr.GetRandomChannel()

		messages, _ := mh.tc.GetChannelMessages(token, channel.Id.String())

		for i := range messages {
			ruby := mh.r.GetReading(messages[i].GetContent())
			userName, _ := mh.ur.GetUserNameById(messages[i].UserId)
			content := messages[i].GetContent()
			if !checkLength(content) {
				continue
			}

			isIka := checkIkaShikaMeka(ruby, ikaRegexp)
			isShika := checkIkaShikaMeka(ruby, shikaRegexp)
			isMeka := checkIkaShikaMeka(content, mekaRegexp)

			mes := domain.Message{
				User:      userName,
				MessageId: uuid.MustParse(messages[i].Id),
				Channel:   channel.Name,
				Content:   content,
				CreatedAt: messages[i].GetCreatedAt(),
				Ika:       isIka,
				Shika:     isShika,
				Meka:      isMeka,
			}

			select {
			case <-ctx.Done():
				return
			case ch <- mes:
			}
		}
	}
}
