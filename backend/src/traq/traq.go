package traq

import (
	"context"

	"github.com/google/uuid"
	gotraq "github.com/traPtitech/go-traq"
)

type TraqClient interface {
	GetMe(string) (*User, error)
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
