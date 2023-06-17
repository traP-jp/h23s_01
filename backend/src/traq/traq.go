package traq

import (
	"context"

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

	var parentId string
	for _, ch := range chanList.Public {
		if ch.Name == parentChannelName {
			parentId = ch.Name
			break
		}
	}

	childChannels := []domain.Channel{}
	for _, ch := range chanList.Public {
		if *ch.ParentId.Get() == parentId {
			childChannels = append(childChannels, domain.Channel{Id: uuid.MustParse(ch.Id), Name: ch.Name})
		}
	}

	return childChannels, nil
}
