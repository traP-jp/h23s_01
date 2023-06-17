package traq

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/traP-jp/h23s_01/backend/src/domain"
	gotraq "github.com/traPtitech/go-traq"
)

type TraqClient interface {
	GetMe(string) (*User, error)
	GetAllChannels(token, paretChannelName string) ([]domain.Channel, error)
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
