package handler

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
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

type messagesSliceWithLock struct {
	mu          sync.Mutex
	messagesHit []domain.Message
	messagesNot []domain.Message
}

type getMessagesGoroutine struct {
	mh   *messageHandler
	mess *messagesSliceWithLock
	ch   chan int
}

const GOROUTINE_NUMBER = 50

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

	mg := &getMessagesGoroutine{
		mh: mh,
		mess: &messagesSliceWithLock{
			messagesHit: make([]domain.Message, 0, count),
			messagesNot: make([]domain.Message, 0, count),
		},
		ch: make(chan int),
	}

	token, err := getToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	for {
		log.Println(1)
		for i := 0; i < GOROUTINE_NUMBER; i++ {
			go mg.getMessages(token)
		}
		chanCount := 0
		for range mg.ch {
			chanCount++
			log.Printf("len: %d", len(mg.mess.messagesHit))
			if chanCount == GOROUTINE_NUMBER || len(mg.mess.messagesHit) >= count*3/10 {
				break
			}
		}

		mg.mess.mu.Lock()
		hitCount := len(mg.mess.messagesHit)
		notHitCount := len(mg.mess.messagesNot)
		if hitCount+notHitCount < count {
			continue
		}
		if hitCount >= count*3/10 && hitCount <= count*4/10 {
			mg.mess.mu.Unlock()
			break
		} else if len(mg.mess.messagesHit) > count*4/10 {
			mg.mess.messagesHit = mg.mess.messagesHit[:count*4/10]
			mg.mess.mu.Unlock()
			break
		}
		mg.mess.mu.Unlock()
	}

	mg.mess.mu.Lock()
	messagesRes := make([]domain.Message, 0, count)
	messagesRes = append(messagesRes, mg.mess.messagesHit...)
	messagesRes = append(messagesRes, mg.mess.messagesNot...)
	mg.mess.mu.Unlock()

	return c.JSON(http.StatusOK, getMessagesRes{
		Count:    count,
		Messages: messagesRes[:count],
	})
}

const LENGTH_LIMIT = 50

var (
	ikaRegexp   = regexp.MustCompile(`(い|イ|ｲ)(か|カ|ｶ)`)
	shikaRegexp = regexp.MustCompile(`(し|シ|ｼ)(か|カ|ｶ)`)
	mekaRegexp  = regexp.MustCompile(`(め|メ|ﾒ)(か|カ|ｶ)`)
)

func checkIkaShikaMeka(content string, reg *regexp.Regexp) bool {
	return reg.MatchString(content)
}

func checkLength(content string) bool {
	return utf8.RuneCountInString(content) <= LENGTH_LIMIT
}

func (mg *getMessagesGoroutine) getMessages(token string) {
	channel, _ := mg.mh.cr.GetRandomChannel()

	messages, _ := mg.mh.tc.GetChannelMessages(token, channel.Id.String())

	for i := range messages {
		content := messages[i].GetContent()
		if !checkLength(content) {
			continue
		}
		ruby := mg.mh.r.GetReading(messages[i].GetContent())
		userName, _ := mg.mh.ur.GetUserNameById(messages[i].UserId)

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

		mg.mess.mu.Lock()
		if isIka || isShika || isMeka {
			mg.mess.messagesHit = append(mg.mess.messagesHit, mes)
		} else {
			mg.mess.messagesNot = append(mg.mess.messagesNot, mes)
		}
		mg.mess.mu.Unlock()
	}

	mg.ch <- 1
}
