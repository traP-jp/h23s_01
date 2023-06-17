package traq

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/traP-jp/h23s_01/backend/src/domain"
	gotraq "github.com/traPtitech/go-traq"
)

type TraqClient interface {
	GetMe(string) (*User, error)
	GetAllChannels(token, paretChannelName string) ([]domain.Channel, error)
	GetChannelMessages(token, channelId string) ([]gotraq.Message, error)
	GetAllUsers(token string) ([]domain.User, error)
}

type traqClient struct {
	client *gotraq.APIClient
}

type User struct {
	Id   uuid.UUID
	Name string
}

func NewTraqClient(client *gotraq.APIClient) TraqClient {
	return &traqClient{
		client: client,
	}
}

func (tc *traqClient) GetMe(token string) (*User, error) {
	meDetail, _, err := tc.client.MeApi.GetMe(context.WithValue(context.Background(), gotraq.ContextAccessToken, token)).Execute()
	if err != nil {
		return nil, err
	}

	me := &User{
		Id:   uuid.MustParse(meDetail.Id),
		Name: meDetail.Name,
	}
	return me, nil
}

func (tc *traqClient) GetAllChannels(token, parentChannelName string) ([]domain.Channel, error) {
	chanList, _, err := tc.client.ChannelApi.GetChannels(context.WithValue(context.Background(), gotraq.ContextAccessToken, token)).Execute()
	if err != nil {
		return nil, err
	}

	allChannels := map[string]gotraq.Channel{}
	var children []string
	dir := strings.Split(parentChannelName, "/")
	for _, ch := range chanList.Public {
		allChannels[ch.Id] = ch
		if ch.Name == dir[0] && ch.ParentId.Get() == nil {
			children = ch.Children
		}
	}

	dir = dir[1:]
	var childChannels []domain.Channel
	for _, d := range dir {
		for _, c := range children {
			if allChannels[c].Name == d {
				children = allChannels[c].Children
				break
			}
		}
	}

	for _, c := range children {
		uuid := uuid.MustParse(allChannels[c].Id)
		name := allChannels[c].Name
		childChannels = append(childChannels, domain.Channel{Id: uuid, Name: name})
	}

	if len(childChannels) == 0 {
		return nil, ErrNoChannel
	}

	return childChannels, nil
}

func (tc *traqClient) GetChannelMessages(token, channelId string) ([]gotraq.Message, error) {
	messages, _, err := tc.client.MessageApi.
		SearchMessages(context.WithValue(context.Background(), gotraq.ContextAccessToken, token)).
		Limit(1).
		Bot(false).
		In(channelId).
		Execute()
	if err != nil {
		return nil, err
	}

	var offset int32
	totalCount := messages.GetTotalHits()
	if totalCount <= 100 {
		offset = 0
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Seed(time.Now().UnixNano())
		offset = rand.Int31n(int32(totalCount) - 100)
	}

	messages, _, err = tc.client.MessageApi.
		SearchMessages(context.WithValue(context.Background(), gotraq.ContextAccessToken, token)).
		Limit(100).
		Offset(offset).
		In(channelId).
		Bot(false).
		Execute()
	if err != nil {
		return nil, err
	}

	return messages.GetHits(), nil
}

func (tc *traqClient) GetAllUsers(token string) ([]domain.User, error) {
	userList, _, err := tc.client.UserApi.GetUsers(context.WithValue(context.Background(), gotraq.ContextAccessToken, token)).IncludeSuspended(true).Execute()
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, user := range userList {
		if !user.Bot {
			uuid := uuid.MustParse(user.Id)
			name := user.Name
			users = append(users, domain.User{Id: uuid, Name: name})
		}
	}

	return users, nil
}
