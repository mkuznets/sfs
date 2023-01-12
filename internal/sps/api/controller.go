package api

import (
	"context"
	"mkuznets.com/go/sps/internal/types"
	"time"
)

type Controller interface {
	GetChannel(ctx context.Context, id string) (*ChannelResponse, error)
	CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error)
	ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error)
}

type controllerImpl struct {
	store Store
}

func NewController(store Store) Controller {
	return &controllerImpl{store: store}
}

// GetChannel godoc
//
//	@Summary	Get channel by ID
//	@Tags		Channels
//	@Produce	json
//	@Param		id	path		string	true	"Channel ID"
//	@Success	200	{array}		ChannelResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/channels/{id} [get]
func (c *controllerImpl) GetChannel(ctx context.Context, id string) (*ChannelResponse, error) {
	model, err := c.store.GetChannel(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &ChannelResponse{
		Id:          model.Id,
		Title:       model.Title,
		Link:        model.Link,
		Authors:     model.Authors,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	return response, nil
}

func (c *controllerImpl) CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error) {
	model := &Channel{
		Id:          RandomChannelId(),
		UserId:      userId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		CreatedAt:   types.NewTime(time.Now()),
		UpdatedAt:   types.NewTime(time.Now()),
		DeletedAt:   nil,
	}

	if err := c.store.CreateChannel(ctx, model); err != nil {
		return nil, err
	}

	response := &ChannelResponse{
		Id:          model.Id,
		Title:       model.Title,
		Link:        model.Link,
		Authors:     model.Authors,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	return response, nil
}

func (c *controllerImpl) ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error) {
	channels, err := c.store.ListChannels(ctx, userId)
	if err != nil {
		return nil, err
	}

	response := make([]*ChannelResponse, 0)
	for _, channel := range channels {
		response = append(response, &ChannelResponse{
			Id:          channel.Id,
			Title:       channel.Title,
			Link:        channel.Link,
			Authors:     channel.Authors,
			Description: channel.Description,
			CreatedAt:   channel.CreatedAt,
			UpdatedAt:   channel.UpdatedAt,
		})
	}

	return response, nil
}
