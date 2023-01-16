package api

import (
	"context"
	"encoding/xml"
	"io"
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/store"
	"mkuznets.com/go/sps/internal/types"
	"os"
)

type Controller interface {
	GetChannel(ctx context.Context, id string) (*ChannelResponse, error)
	GetChannelFeed(ctx context.Context, channelId string) ([]byte, error)
	CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error)
	ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error)
	UploadFile(ctx context.Context, userId string, f io.ReadSeeker) (*UploadResponse, error)
	CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*EpisodeResponse, error)
	ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error)
}

type controllerImpl struct {
	uploader  Uploader
	store     store.Store
	idService IdService
}

func NewController(store store.Store, uploader Uploader, idService IdService) Controller {
	return &controllerImpl{
		store:     store,
		uploader:  uploader,
		idService: idService,
	}
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

func (c *controllerImpl) GetChannelFeed(ctx context.Context, channelId string) ([]byte, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	return channel.Feed.Content, nil
}

func (c *controllerImpl) CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error) {
	model := &store.Channel{
		Id:          c.idService.Channel(),
		UserId:      userId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		CreatedAt:   types.NewTimeNow(),
		UpdatedAt:   types.NewTimeNow(),
	}

	podcast := ChannelToPodcast(model, nil)
	feed, err := xml.Marshal(podcast)
	if err != nil {
		return nil, err
	}

	model.Feed = store.Feed{
		Content:     feed,
		Url:         "",
		PublishedAt: types.NewTimeNow(),
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

func (c *controllerImpl) UploadFile(ctx context.Context, userId string, f io.ReadSeeker) (*UploadResponse, error) {
	tmpFile, err := os.CreateTemp("", "tmpUpload")
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if err != nil {
		return nil, err
	}

	info, err := c.uploader.Upload(ctx, f)
	if err != nil {
		return nil, err
	}

	model := &store.File{
		Id:          c.idService.File(),
		UserId:      userId,
		Url:         info.Url,
		Size:        info.Size,
		ContentType: info.ContentType,
		CreatedAt:   types.NewTimeNow(),
		UpdatedAt:   types.NewTimeNow(),
	}
	if err := c.store.CreateFile(ctx, model); err != nil {
		return nil, err
	}

	return &UploadResponse{
		Id: model.Id,
	}, nil
}

func (c *controllerImpl) CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*EpisodeResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, herror.NotFound("channel not found")
	}

	file, err := c.store.GetFile(ctx, r.FileId)
	if err != nil {
		return nil, err
	}
	if file.UserId != userId {
		return nil, herror.NotFound("file not found")
	}

	model := &store.Episode{
		Id:          c.idService.Episode(),
		ChannelId:   channelId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		FileId:      r.FileId,
		CreatedAt:   types.NewTimeNow(),
		UpdatedAt:   types.NewTimeNow(),
	}

	if err := c.store.CreateEpisode(ctx, model); err != nil {
		return nil, err
	}

	response := &EpisodeResponse{
		Id: model.Id,
		File: &FileResponse{
			Id:          file.Id,
			Url:         file.Url,
			Size:        file.Size,
			ContentType: file.ContentType,
		},
		ChannelId:   model.ChannelId,
		Title:       model.Title,
		Link:        model.Link,
		Authors:     model.Authors,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	return response, nil
}

func (c *controllerImpl) ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, herror.NotFound("channel not found")
	}

	episodes, err := c.store.ListEpisodesWithFiles(ctx, channelId)
	if err != nil {
		return nil, err
	}

	response := make([]*EpisodeResponse, 0)
	for _, episode := range episodes {
		response = append(response, &EpisodeResponse{
			Id: episode.Id,
			File: &FileResponse{
				Id:          episode.File.Id,
				Url:         episode.File.Url,
				Size:        episode.File.Size,
				ContentType: episode.File.ContentType,
			},
			ChannelId:   episode.ChannelId,
			Title:       episode.Title,
			Link:        episode.Link,
			Authors:     episode.Authors,
			Description: episode.Description,
			CreatedAt:   episode.CreatedAt,
			UpdatedAt:   episode.UpdatedAt,
		})
	}

	return response, nil
}
