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

// messageに関する並行処理を担当する構造体
type messageWorker struct {
	mh *messageHandler
	ch chan Result
}

// messageを取得したときの結果
type Result struct {
	message domain.Message
	err     error
}

func newMessageWorker(mh *messageHandler) *messageWorker {
	return &messageWorker{
		mh: mh,
		ch: make(chan Result),
	}
}

// goroutineの数
const GOROUTINE_NUMBER = 30

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

	mw := newMessageWorker(mh)
	for i := 0; i < GOROUTINE_NUMBER; i++ {
		go mw.getMessage(ctx)
	}

	for {
		r := <-mw.ch
		if r.err != nil {
			cancel()
			// 本来ならエラーの種類に合わせて適切なレスポンスを返す
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		mes := r.message
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

// メッセージを取得して送信
// tokenはctxに含める
func (mw *messageWorker) getMessage(ctx context.Context) {
	token := ctx.Value(key).(string)
	for {
		channel, err := mw.mh.cr.GetRandomChannel()
		if err != nil {
			mw.ch <- Result{err: err}
			return
		}

		messages, err := mw.mh.tc.GetChannelMessages(token, channel.Id.String())
		if err != nil {
			mw.ch <- Result{err: err}
			return
		}

		for i := range messages {
			ruby := mw.mh.r.GetReading(messages[i].GetContent())
			userName, err := mw.mh.ur.GetUserNameById(messages[i].UserId)
			if err != nil {
				mw.ch <- Result{err: err}
				return
			}
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
			case mw.ch <- Result{message: mes}:
			}
		}
	}
}
